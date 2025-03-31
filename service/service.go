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
	"os"
	"os/signal"
	"syscall"
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
}

func (s *service) Stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGTERM, /*  Согласно всякой документации именно он должен останавливать прогу, но на деле его мы не находим. Оставил его просто на всякий случай  */
		syscall.SIGINT,  /*  Останавливает прогу когда она запущена из терминала и останавливается через CTRL+C  */
		syscall.SIGQUIT, /*  Останавливает демона systemd  */
	)

	<-quit

	s.grpcServer.GracefulStop()
	s.logger.Info("Graceful Stop", nil)
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
