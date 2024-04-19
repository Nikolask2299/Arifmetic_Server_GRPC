package app

import (
	"fmt"
	"net"
	"services/interal/serverauth"

	"google.golang.org/grpc"
)


type AuthApp struct {
	gRPCServer *grpc.Server
	port int
}


func NewAuthApp(authserv *serverauth.ServerAuth, port int) *AuthApp {
	gRPCServer := grpc.NewServer()
	serverauth.Register(gRPCServer, authserv)
	return &AuthApp{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *DataApp) MustRun() {

	if err := app.Run(); err != nil {
		panic(err)
	} else {
		fmt.Println("ServerData is running OK")
	}
}

func (app *DataApp) Run() error {
	host := "localhost"

	addr := fmt.Sprintf("%s:%d", host, app.port)
	
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if err := app.gRPCServer.Serve(lis); err != nil {
		return err
	}
	
	return nil
}

func (app *DataApp) Stop() {
	app.gRPCServer.GracefulStop()
}
