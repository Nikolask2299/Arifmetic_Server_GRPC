package serverauth

import (
	"context"
	"fmt"
	serv1 "protobuf/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



type AuthClient struct {
	timeout time.Duration
	port string
}

func NewAuthClient(timeout time.Duration, port string) *AuthClient {
	return &AuthClient{
		timeout: timeout,
		port: port,
	}
}

func (c *AuthClient) NewAuthClientConnection() (context.Context, serv1.AuthorizationServiceClient) {

	addr := fmt.Sprintf("localhost:%s", c.port)

	ctx, _ := context.WithTimeout(context.Background(), c.timeout)

	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := serv1.NewAuthorizationServiceClient(conn)

	return ctx, client
}