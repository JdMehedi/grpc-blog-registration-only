package postgres

import (
	"context"
	"grpc-blog/blog/storage"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Post
		want    int64
		wantErr bool
	}{
		{
			name:"CREATE_Post_SUCCESS",
			in: storage.Post{
				Title: "This is title",
				Description: "This is description",
			},
			want: 1,
		},

		{
			name:"CREATE_Post_SUCCESS",
			in: storage.Post{
				Title: "This is title 2",
				Description: "This is description 2",
			},
			want: 2,
		},
		

		
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Create(context.Background(), tt.in)

			log.Printf("%#v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErrGet %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, get %v", got, tt.want)
			}
		})
	}
}

