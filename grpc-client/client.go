package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"

	mainapipb "simplegrpcclient/proto/gen"
	farewellpb "simplegrpcclient/proto/gen/farewell"
)

func main() {
	cert := "certs/cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Failed to load cert:", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalln("Error connecting:", err)
	}

	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)

	clientGreet := mainapipb.NewGreetingClient(conn)

	fwClient := farewellpb.NewAufWiedersehenClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	res, err := client.Add(ctx, req)

	if err != nil {
		log.Fatalln("Error connecting", err)
	}

	reqGreet := &mainapipb.HelloRequest{
		Name: "John",
	}

	resGreet, err := clientGreet.Greet(ctx, reqGreet)
	if err != nil {
		log.Fatalln("Could not greet:", err)
	}

	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "John",
	}

	resFw, err := fwClient.BidGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatalln("Could not say bye bye:", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("---------------------------------------------")
	log.Println("Greeting message:", resGreet.Message)
	log.Println("---------------------------------------------")
	log.Println("Bye message:", resFw.Message)
}
