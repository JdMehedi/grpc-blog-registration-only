package post

import (
	"context"
	"grpc-blog/blog/storage"
	tgv "grpc-blog/gunk/v1/post"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)


func (s *Svc) CreatePost(ctx context.Context, req *tgv.CreatePostRequest) (*tgv.CreatePostResponse, error){
	//Need to Validet post
	
	post:=storage.Post{
		ID: req.GetPost().ID,
		Title: req.GetPost().Title,
		Description: req.GetPost().Description,
	}

	id, err:= s.core.Create(context.Background(), post)
	if err!=nil{
		return nil, status.Error(codes.Internal,"Failed to create post:")
	}

	return &tgv.CreatePostResponse{
		ID: id,
	},nil
}