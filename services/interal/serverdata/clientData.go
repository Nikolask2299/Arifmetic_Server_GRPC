package serverdata

import (
	"context"
	"fmt"
	serv1 "protobuf/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataClient struct {
	timeout time.Duration
	port string
}

func NewDataClient(timeout time.Duration, port string) *DataClient {
	return &DataClient{
		timeout: timeout,
		port: port,
	}
}

func (c *DataClient) NewDataClientConnection() (context.Context, serv1.DataBaseServiceClient) {

	addr := fmt.Sprintf("localhost:%s", c.port)

	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := serv1.NewDataBaseServiceClient(conn)

	return ctx, client
}



