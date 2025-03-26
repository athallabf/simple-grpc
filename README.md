# Simple gRPC Project

A basic gRPC implementation in Go demonstrating client-server communication using Protocol Buffers.

## Project Structure

```
.
├── client/
│   └── main.go
├── server/
│   └── main.go
├── proto/
│   └── greet.proto
├── greetpb/
│   ├── greet.pb.go
│   └── greet_grpc.pb.go
├── go.mod
└── go.sum
```

## Features

- Simple gRPC server implementation
- Unary RPC call with `SayHello` method
- Protocol Buffer message definitions
- Basic client implementation

## Prerequisites

- Go 1.24.1 or later
- Protocol Buffers compiler (protoc)
- Go gRPC plugins

## Installation

1. Clone the repository:

```
git clone https://github.com/athallabf/simple-grpc.git
cd simple-grpc
```

2. Install dependencies:

```
go mod tidy
```

## Generate Protocol Buffer Code

To regenerate the Protocol Buffer code, run:

```
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

## Running the Application

1. Start the server:

```
go run server/main.go
```

2. In a separate terminal, run the client:

```
go run client/main.go
```

## Service Definition

The service is defined in the proto file with a simple greeting service:

```
service GreetService {
  rpc SayHello(GreetRequest) returns (GreetResponse);
}
```

## Dependencies

- google.golang.org/grpc v1.71.0
- google.golang.org/protobuf v1.36.6
- Other standard Go packages

## Author

Athalla
