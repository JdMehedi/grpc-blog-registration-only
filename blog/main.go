package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-blog/blog/services/post"
	"grpc-blog/blog/storage/postgres"

	tgv "grpc-blog/gunk/v1/post"
	tcp "grpc-blog/blog/core/post"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("Failed t listen: %s", err)
	}

	grpcServer := grpc.NewServer()

	store, err :=postgres.NewStorage("dbstring")
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	cs := tcp.NewCoreSvc(store)

	s := post.NewSvc(cs)

	tgv.RegisterPostServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {

		log.Fatalf("Failed to serve: %s", err)
	}

}
