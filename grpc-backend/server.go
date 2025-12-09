package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "simplegrpcserver/proto/gen"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReponse, error) {
	return &pb.AddReponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {

	cert := "certs/cert.pem"
	key := "certs/key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Error loading cert:", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}

}
