package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql SqlStorage) CreateUser(ctx context.Context, data *entity.User) error {
	if err := sql.db.Create(&data).Error; err != nil {
		return entity.ErrorDB(err)
	}
	return nil
}
