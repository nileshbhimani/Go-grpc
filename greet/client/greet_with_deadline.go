package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Niles",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {

			if e.Code() == codes.InvalidArgument {
				log.Println("Deadline exceeded!: ")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}

		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

	}
	log.Printf("GreetWithDeadlines : %s\n", res.Result)
}
