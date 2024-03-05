package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Println("doGreetEveryOne was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement "},
		{FirstName: "Marie"},
		{FirstName: "Tests"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err != nil {
				log.Printf("Error while Receiving : %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)

		}

		close(waitc)
	}()

	<-waitc

}
