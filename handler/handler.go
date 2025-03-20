package handler

import (
	"authService/authStorage"
	proto "authService/proto"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type handler struct {
	proto.UnimplementedAuthServer
}

func New() *handler {
	return &handler{}
}

const jwtSecretKey = "very-secret-key"

func (s *handler) Registration(ctx context.Context, in *proto.RegistrationRequest) (*proto.RegistrationResponse, error) {
	if _, ok := authStorage.AuthStorage[in.Email]; ok {
		return &proto.RegistrationResponse{
			Error: "User already registered",
		}, nil
	}

	authStorage.AuthStorage[in.Email] = &authStorage.User{
		Name:  in.Name,
		Pass:  in.Password,
		Email: in.Email,
	}

	return &proto.RegistrationResponse{
		Message: "Success registered",
	}, nil
}

func (s *handler) Authentication(ctx context.Context, in *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	if _, ok := authStorage.AuthStorage[in.Email]; !ok {
		return &proto.AuthenticationResponse{
			Error: "такой почтовый ящик не зарегистрирован в системе",
		}, nil
		//return nil, errors.New("такой почтовый ящик не зарегистрирован в системе")
	}

	if authStorage.AuthStorage[in.Email].Name != in.Name {
		return &proto.AuthenticationResponse{
			Error: "имя неверное",
		}, nil
		//return nil, errors.New("имя неверное")
	}

	if authStorage.AuthStorage[in.Email].Pass != in.Password {
		return &proto.AuthenticationResponse{
			Error: "пароль неверный",
		}, nil
		//return nil, errors.New("пароль неверный")
	}

	payload := jwt.MapClaims{
		"sub": in.Email + "_" + in.Password,       //
		"exp": time.Now().Add(time.Minute).Unix(), //время жизни
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return &proto.AuthenticationResponse{
			Error:       "ошибка подписи",
			AccessToken: signedToken,
		}, nil
		//return nil, errors.New("ошибка подписи")
	}

	authStorage.AuthStorage[in.Email].Token = signedToken

	return &proto.AuthenticationResponse{
		Message:     "Success authenticated",
		AccessToken: signedToken,
	}, nil
}

func (s *handler) Authorization(ctx context.Context, in *proto.AuthorizationRequest) (*proto.AuthorizationResponse, error) {
	//расшифровываем токен
	//достаем пэйлоад
	//по пэйлоаду получаем эмеил
	//сравниваем токен в хранилище с полученным

	//fmt.Println("получил структуру в авторизации", in.AccessToken)

	in.AccessToken = strings.ReplaceAll(in.AccessToken, "Bearer ", "")
	sToken := strings.Split(in.AccessToken, ".")

	type payload struct {
		Sub string `json:"sub"`
		Exp int64  `json:"exp"`
	}

	//по какой-то причине декодеровщик обрезает последнюю фигурную скобку
	payload1, _ := base64.StdEncoding.DecodeString(sToken[1])
	//fmt.Println("payload1", string(payload1), sToken[1], payload1)
	payload1 = append(payload1, 125)

	var p payload

	err := json.Unmarshal(payload1, &p)
	if err != nil {
		return &proto.AuthorizationResponse{
			Error: fmt.Sprint("Ошибка сервиса авторизации", err),
		}, err
		//log.Fatal(err)
	}

	//fmt.Println(p)

	sub := strings.Split(p.Sub, "_")

	//fmt.Println(sub)

	if authStorage.AuthStorage[sub[0]].Token != in.AccessToken {
		return &proto.AuthorizationResponse{
			Error: "невалидный токен",
		}, nil
	}

	ti := time.Now()

	if ti.After(time.Unix(p.Exp, 0)) == true {
		return &proto.AuthorizationResponse{
			Error: "токен просрочен",
		}, nil
	}

	return &proto.AuthorizationResponse{
		Message: "Success authorized",
	}, nil
}

func (s *handler) Logout(ctx context.Context, in *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	return &proto.LogoutResponse{}, nil

}
