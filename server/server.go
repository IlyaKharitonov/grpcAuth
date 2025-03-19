package server

import (
	proto "authService/proto"
	"context"
)

type server struct {
	proto.UnimplementedAuthServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Registration(ctx context.Context, in *proto.RegistrationRequest) (*proto.RegistrationResponse, error) {
	return &proto.RegistrationResponse{}, nil
}

func (s *server) Authentication(ctx context.Context, in *proto.AuthenticationRequest) (*proto.AuthenticationResponse, error) {
	return &proto.AuthenticationResponse{}, nil
}

func (s *server) Authorization(ctx context.Context, in *proto.AuthorizationRequest) (*proto.AuthorizationResponse, error) {
	return &proto.AuthorizationResponse{}, nil
}

func (s *server) Logout(ctx context.Context, in *proto.LogoutRequest) (*proto.LogoutResponse, error) {
	return &proto.LogoutResponse{}, nil

}
