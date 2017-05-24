package main

import (
	"gomad"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8182")
	if err != nil {
		grpclog.Fatalf("broken: %s", err)
	}

	grpcServer := grpc.NewServer()
	gomad.RegisterGomadServiceServer(grpcServer, gomad.NewGomadServiceServer())
	grpcServer.Serve(lis)
}
