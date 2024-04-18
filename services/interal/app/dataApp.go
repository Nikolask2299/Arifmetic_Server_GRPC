package app

import (
	"fmt"
	"net"
	"services/interal/serverdata"
	"google.golang.org/grpc"
)

type DataApp struct {
	gRPCServer *grpc.Server
	port       int
}

func NewDataApp(dateserv *serverdata.ServerData, port int) *DataApp {
	gRPCServer := grpc.NewServer()
	serverdata.Register(gRPCServer, dateserv)
	return &DataApp{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (app *DataApp) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
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
