package storage

import (
	"context"
	"indetity-service/dto"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) UpdatePermission(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoPermissionUpdate) error {
	if err := sql.db.Where("name = ?", cond["name"]).Updates(dataUpdate).Error; err != nil {
		return entity.CantUpdatePermission(err)
	}

	return nil
}
