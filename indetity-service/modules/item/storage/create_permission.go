package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) CreatePermission(ctx context.Context, en *entity.Permission) error {
	if err := sql.db.Create(&en).Error; err != nil {
		return entity.CantCreatePermission(err)
	}

	return nil
}
