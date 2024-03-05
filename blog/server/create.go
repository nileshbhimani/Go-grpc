package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create blog was invoked with : %v\n", in)

	data := BlogItem{
		AutherId: in.AutherId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := colletion.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error : %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to Oid",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil

}
