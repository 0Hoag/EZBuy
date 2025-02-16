package storage

import (
	"context"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) DeleteUserById(ctx context.Context, cond map[string]interface{}) error {
	if err := sql.db.Table(entity.User{}.TableName()).Where("user_id = ?", cond["user_id"]).Delete(nil).Error; err != nil {
		return entity.CantDeleteUser(err)
	}

	return nil
}
