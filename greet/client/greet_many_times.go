package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	// res, err := c.Greet(context.Background(), &pb.GreetRequest{
	// 	FirstName: "Clement",
	// })

	// if err != nil {
	// 	log.Fatalf("Could not Greet %v\n", err)
	// }

	// log.Printf("Greeting %s\n", res.Result)

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not GreetManyTimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		// server connection is closed
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading %v\n", err)
		}

		log.Printf("GreetManyTimes %s\n", msg.Result)

	}

}
