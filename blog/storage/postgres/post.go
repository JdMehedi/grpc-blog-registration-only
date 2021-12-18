package postgres

import (
	"context"
	"grpc-blog/blog/storage"
	"log"
)
const insertTodo = `
	INSERT INTO posts(
		title,
		description
	) VALUES(
		:title,
		:description
	)RETURNING id;
`

func (s Storage) Create(ctx context.Context,t storage.Post)(int64, error){
	stmt, err := s.db.PrepareNamed(insertTodo)
	if err != nil {
		return 0, err
	}
	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	log.Println("Todo ID: ", id)
	return id, nil
}