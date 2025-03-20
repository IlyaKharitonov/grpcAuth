package service

import (
	"authService/config"
	"authService/handler"
	proto "authService/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) Run() {
	//начинает слушать на указанном
	//создает grpc сервер
	//регистрирует обработчики к которым будем обращаться по рпс
	config.ParseConfig()
	fmt.Println("Распарсил конфиг")

	lis, err := net.Listen("tcp", fmt.Sprintf(config.GetConfig().Server.Host+":"+config.GetConfig().Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Создал слушателя")

	server := grpc.NewServer()

	proto.RegisterAuthServer(
		server,
		handler.New(),
	)

	fmt.Println("Зарегистрировал обработчик")

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Запустил сервис")
	//server.GracefulStop()
}

func (s *service) Stop() {

}
