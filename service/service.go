package service

import (
	"authService/config"
	proto "authService/proto"
	"authService/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Run() {
	//начинает слушать на указанном
	//создает grpc сервер
	//регистрирует обработчики к которым будем обращаться по рпс
	config.ParseConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf(config.GetConfig().Server.Host+":"+config.GetConfig().Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterAuthServer(
		grpc.NewServer(),
		server.NewServer())

}

func (s *service) Stop() {

}
