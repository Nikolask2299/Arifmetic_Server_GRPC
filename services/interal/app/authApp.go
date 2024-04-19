package app

import (
	"fmt"
	"net"
	"services/interal/serverauth"

	"google.golang.org/grpc"
)


type AuthApp struct {
	gRPCServer *grpc.Server
	port string
}


func NewAuthApp(authserv *serverauth.ServerAuth, port string) *AuthApp {
	gRPCServer := grpc.NewServer()
	serverauth.Register(gRPCServer, authserv)
	return &AuthApp{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *AuthApp) MustRun() {
	
	if err := app.Run(); err != nil {
		panic(err)
	} 
}

func (app *AuthApp) Run() error {
	host := "localhost"

	addr := fmt.Sprintf("%s:%s", host, app.port)
	
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if err := app.gRPCServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (app *AuthApp) Stop() {
	app.gRPCServer.GracefulStop()
}
