package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) CreateRole(ctx context.Context, role *entity.Role) error {
	if err := sql.db.Create(&role).Error; err != nil {
		return entity.CantCreateRole(err)
	}

	return nil
}
