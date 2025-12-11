package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "simplegrpcserver/proto/gen"
	farewell "simplegrpcserver/proto/gen/farewell"
	farewellpb "simplegrpcserver/proto/gen/farewell"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreetingServer
	farewell.UnimplementedAufWiedersehenServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReponse, error) {
	return &pb.AddReponse{
		Sum: req.A + req.B,
	}, nil
}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s. Nice to meet you", req.Name),
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, req *farewell.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	return &farewellpb.GoodByeResponse{
		Message: fmt.Sprintf("Goodbye %s. Nice to meet you", req.Name),
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
	pb.RegisterGreetingServer(grpcServer, &server{})
	farewell.RegisterAufWiedersehenServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to serve:", err)
	}

}
