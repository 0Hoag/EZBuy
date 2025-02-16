package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type SqlStorage struct {
	db *mongo.Client
}

func NewSqlStore(db *mongo.Client) *SqlStorage {
	return &SqlStorage{db: db}
}
