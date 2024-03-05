package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func UpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("==== UpdateBlog was invoked =====")

	newBlog := &pb.Blog{
		Id:       id,
		AutherId: "Not Clement",
		Title:    "A new title",
		Content:  "Content of the frist with some awsom additon",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating:%v\n", err)
	}

	log.Println("Blog was updated!")

}
