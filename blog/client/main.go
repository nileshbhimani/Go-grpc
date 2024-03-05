package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect on: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	doReadBlog(c, id) // valid
	UpdateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
	// doReadBlog(c, "aNonExistingId")

	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, 10)
}
