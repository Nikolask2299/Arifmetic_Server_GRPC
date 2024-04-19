package serverauth

import (
	"context"
	serv "protobuf/gen/go"
	"services/interal/service/auth"

	"google.golang.org/grpc"
)

type ServerAuth struct {
	serv.AuthorizationServiceServer
	auth auth.Auth
}

func NewServerAuth(auth auth.Auth) *ServerAuth {
	return &ServerAuth{auth: auth}
}

func Register(s *grpc.Server, sa *ServerAuth) {
	serv.RegisterAuthorizationServiceServer(s, sa)
}

type AuthorizationServiceServer interface {
	Login(context.Context, *serv.LoginRequest) (*serv.LoginResponse, error)
	Register(context.Context, *serv.RegisterRequest) (*serv.RegisterResponse, error)	
}

func (s *ServerAuth) LoginUser(ctx context.Context, req *serv.LoginRequest) (*serv.LoginResponse, error) {
	user := req.GetUserName()
	pass := req.GetPassword()

	res, err := s.auth.Login(ctx, user, pass)
	if err != nil {
		return nil, err
	}

	return &serv.LoginResponse{
		Token: res,
	}, nil			
}

func (s *ServerAuth) RegisterUser(ctx context.Context, req *serv.RegisterRequest) (*serv.RegisterResponse, error) {
	user := req.GetUserName()
	pass := req.GetPassword()

	res, err := s.auth.Register(ctx, user, pass)
	if err != nil {
		return nil, err
	}

	return &serv.RegisterResponse{
		UserId: res,
	}, nil
}
