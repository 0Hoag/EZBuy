package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) DeleteRole(ctx context.Context, cond map[string]interface{}) error {
	err := sql.db.Table(entity.Role{}.TableName()).Where("name = ?", cond["name"]).Delete(nil).Error
	if err != nil {
		return entity.RoleNotFound
	}

	return nil
}
