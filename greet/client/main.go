package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"

		cred, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(cred))
	}

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect on: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryOne(c)
	// doGreetWithDeadline(c, 1*time.Second)
	// doGreetWithDeadline(c, 5*time.Second)
}
