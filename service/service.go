package service

import (
	"authService/config"
	"authService/handler"
	proto "authService/proto"
	"database/sql"
	"fmt"
	"github.com/IlyaKharitonov/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct {
	connDB     *sql.DB
	listener   net.Listener
	grpcServer *grpc.Server

	logger logger.ILogger
}

func New() *service {
	config := config.ParseConfig()

	logger := logger.New(&config.Logger)
	logger.Info("Запустил логгер", nil)

	//add db conn

	listener := NewListener(&config.Server)
	logger.Info(
		"Запустил tcp слушателя",
		nil,
		logger.AddParam("Host", config.Server.Host),
		logger.AddParam("Port", config.Server.Port))

	grpcServer := NewGrpcServer()
	logger.Info("Создал объект GRPC сервера и зарегистрировал обработчики", nil)

	return &service{
		connDB:     nil,
		listener:   listener,
		grpcServer: grpcServer,

		logger: logger,
	}
}

func (s *service) Run() {
	s.logger.Info("Запустил GRPC сервер", nil)

	err := s.grpcServer.Serve(s.listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//server.GracefulStop()
}

func (s *service) Stop() {

}

func NewListener(config *config.ServerConf) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(config.Host+":"+config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}

func NewGrpcServer() *grpc.Server {
	server := grpc.NewServer()

	proto.RegisterAuthServer(
		server,
		handler.New(),
	)

	return server
}
