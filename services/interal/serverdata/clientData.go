package serverdata

import (
	"context"
	"fmt"
	serv1 "protobuf/gen/go"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewDataClientConnection(timeout time.Duration, port int) (context.Context, serv1.DataBaseServiceClient) {

	addr := fmt.Sprintf("localhost:%d", port)

	ctx, _ := context.WithTimeout(context.Background(), timeout)

	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := serv1.NewDataBaseServiceClient(conn)

	return ctx, client
}



