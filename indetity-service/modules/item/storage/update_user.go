package storage

import (
	"context"
	"indetity-service/dto"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoUpdateRequest) error {
	if err := sql.db.Where("user_id = ?", cond["user_id"]).Updates(dataUpdate).Error; err != nil {
		return entity.CantUpdateUser(err)
	}
	return nil
}
