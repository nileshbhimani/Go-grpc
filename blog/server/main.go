package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	pb.BlogServiceServer
}

var addr string = "0.0.0.0:50051"
var colletion *mongo.Collection

func main() {

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@mongo:27017/"))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = client.Connect(context.Background())

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		panic(err)
	}

	colletion = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	// pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
