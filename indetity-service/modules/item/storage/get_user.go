package storage

import (
	"context"
	"gorm.io/gorm"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) GetUserById(ctx context.Context, cond map[string]interface{}) (*entity.User, error) {
	var data entity.User

	if err := sql.db.Preload("Role.Permission").Where("user_id = ?", cond["user_id"]).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.UserNotFound
		}
		return nil, entity.ErrorDB(err)
	}

	return &data, nil
}

func (sql *SqlStorage) GetUserByUserName(ctx context.Context, cond map[string]interface{}) (*entity.User, error) {
	var data entity.User

	if err := sql.db.Preload("Role.Permission").Where("username = ?", cond["username"]).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.UserNotFound
		}
		return nil, modules.ErrDB(err)
	}

	return &data, nil
}

func (sql *SqlStorage) GetAllUser(ctx context.Context, paging *modules.Paging, moreKeys ...string) ([]entity.User, error) {
	var result []entity.User

	paging.Process()

	if err := sql.db.Table(entity.User{}.TableName()).Count(&paging.TotalPage).Error; err != nil {
		return nil, entity.CantGetListUser(err)
	}

	if err := sql.db.Offset((paging.Page - 1) * paging.Size).
		Limit(paging.Size).
		Preload("Role.Permission").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
