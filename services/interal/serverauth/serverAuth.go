package serverauth

import (
	"context"
	serv1 "protobuf/gen/go"

	"google.golang.org/grpc"
)

type ServerAuth struct {
	serv1.AuthorizationServiceServer
}

func NewServerAuth() *ServerAuth {
	return &ServerAuth{}
}

func Register(s *grpc.Server, sa *ServerAuth) {
	serv1.RegisterAuthorizationServiceServer(s, sa)
}

type AuthorizationServiceServer interface {
	Login(context.Context, *serv1.LoginRequest) (*serv1.LoginResponse, error)
	Register(context.Context, *serv1.RegisterRequest) (*serv1.RegisterResponse, error)	
}

func (s *ServerAuth) Login(ctx context.Context, req *serv1.LoginRequest) (*serv1.LoginResponse, error) {

}

func (s *ServerAuth) Register(ctx context.Context, req *serv1.RegisterRequest) (*serv1.RegisterResponse, error) {

}
