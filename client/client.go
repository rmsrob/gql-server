package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcClient(addr, name string) string {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to dial: %v", err)
	}

	defer conn.Close()
	c := NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Printf("Failed to call: %v", err)
	}

	return r.GetMessage()
}
