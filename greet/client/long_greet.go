package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement "},
		{FirstName: "Marie"},
		{FirstName: "Tests"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while handling error %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
