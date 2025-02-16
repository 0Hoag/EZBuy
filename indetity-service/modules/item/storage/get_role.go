package storage

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) GetRoleById(ctx context.Context, cond map[string]interface{}) (*entity.Role, error) {
	var data entity.Role

	if err := sql.db.Preload("Permission").Where("name = ?", cond["name"]).First(&data).Error; err != nil {
		return nil, entity.RoleNotFound
	}

	return &data, nil
}

func (sql *SqlStorage) GetAllRole(ctx context.Context, paging *modules.Paging) ([]entity.Role, error) {
	var result []entity.Role

	paging.Process()

	if err := sql.db.Table(entity.Role{}.TableName()).Count(&paging.TotalPage).Error; err != nil {
		return nil, entity.CantGetListRole(err)
	}

	if err := sql.db.Offset((paging.Page - 1) * paging.Size).
		Limit(paging.Size).
		Preload("Permission").
		Find(&result).Error; err != nil {
		return nil, entity.CantGetListRole(err)
	}

	return result, nil
}
