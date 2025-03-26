// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"

// 	"github.com/athallabf/simple-grpc/greetpb"

// 	"google.golang.org/grpc"
// )

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/athallabf/simple-grpc/greetpb"

	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &greetpb.GreetResponse{Message: message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	log.Println("Server is running on port 50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
