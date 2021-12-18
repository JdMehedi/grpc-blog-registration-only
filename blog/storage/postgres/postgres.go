package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)


type Storage struct{
	db *sqlx.DB
}

func NewStorage(dbst string) (*Storage, error){
	db, err := sqlx.Connect("postgres",dbst)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postgres '%s'", dbst)
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return &Storage{db: db}, nil
}