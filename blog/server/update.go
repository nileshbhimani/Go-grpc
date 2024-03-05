package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog was invoked:%v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Can not pasrse ID",
		)
	}

	data := &BlogItem{
		AutherId: in.AutherId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := colletion.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with id",
		)
	}

	return &emptypb.Empty{}, nil
}
