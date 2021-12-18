package post

import (
	"context"

	"grpc-blog/blog/storage"


	tgv "grpc-blog/gunk/v1/post"
)

type PostCoreStore interface{
	Create(context.Context, storage.Post)(int64, error)
}

type Svc struct {
	tgv.UnimplementedPostServiceServer
	core PostCoreStore
}

func NewSvc(s PostCoreStore) *Svc {
	return &Svc{
		core: s,
	}
}


