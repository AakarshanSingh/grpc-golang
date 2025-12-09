package main

import (
	"context"
	"log"
	mainapipb "simplegrpcclient/proto/gen"
	"time"

	"google.golang.org/grpc"
	
)

func main() {
	conn, err := grpc.NewClient("localhost:50051",)

	if err != nil {
		log.Fatalln("Error connecting", err)
	}

	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	req := mainapipb.AddRequest{
		A: 10,
		B: 20,
	}

	res, err := client.Add(ctx, &req)

	if err != nil {
		log.Fatalln("Error connecting", err)
	}

	log.Println("Sum:", res.Sum)

}
