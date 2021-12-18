package post

import (
	"context"
	"grpc-blog/blog/storage"


)
type PoststorageStore interface{
	Create(context.Context, storage.Post)(int64, error)
}

type CoreSvc struct{
	store PoststorageStore
}

func NewCoreSvc(sp PoststorageStore) *CoreSvc{
	return &CoreSvc{
		store: sp,
	}
}

func (cs CoreSvc) Create(ctx context.Context,t storage.Post)(int64, error){

	return cs.store.Create(ctx,t)
}