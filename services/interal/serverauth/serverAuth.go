package serverauth

import (
	"context"
	serv1 "protobuf/gen/go"
	"services/interal/service/auth"

	"google.golang.org/grpc"
)

type ServerAuth struct {
	serv1.AuthorizationServiceServer
	auth auth.Auth
}

func NewServerAuth(auth auth.Auth) *ServerAuth {
	return &ServerAuth{auth: auth}
}

func Register(s *grpc.Server, sa *ServerAuth) {
	serv1.RegisterAuthorizationServiceServer(s, sa)
}

type AuthorizationServiceServer interface {
	Login(context.Context, *serv1.LoginRequest) (*serv1.LoginResponse, error)
	Register(context.Context, *serv1.RegisterRequest) (*serv1.RegisterResponse, error)	
}

func (s *ServerAuth) Login(ctx context.Context, req *serv1.LoginRequest) (*serv1.LoginResponse, error) {
	user := req.GetUserName()
	pass := req.GetPassword()

	res, err := s.auth.Login(ctx, user, pass)
	if err != nil {
		return nil, err
	}

	return &serv1.LoginResponse{
		Token: res,
	}, nil			
}

func (s *ServerAuth) Register(ctx context.Context, req *serv1.RegisterRequest) (*serv1.RegisterResponse, error) {
	user := req.GetUserName()
	pass := req.GetPassword()

	res, err := s.auth.Register(ctx, user, pass)
	if err != nil {
		return nil, err
	}

	return &serv1.RegisterResponse{
		UserId: res,
	}, nil
}
