package storage

import "gorm.io/gorm"

type SqlStorage struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *SqlStorage {
	return &SqlStorage{db: db}
}
