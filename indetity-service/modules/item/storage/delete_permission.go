package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql SqlStorage) DeletePermission(ctx context.Context, cond map[string]interface{}) error {
	err := sql.db.Table(entity.Permission{}.TableName()).Where("name = ?", cond["name"]).Delete(nil).Error
	if err != nil {
		return entity.CantDeletePermission(err)
	}

	return nil
}
