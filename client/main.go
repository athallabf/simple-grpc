package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/athallabf/simple-grpc/greetpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	req := &greetpb.GreetRequest{Name: "Athalla"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	fmt.Println("Response from server:", res.GetMessage())
}
