package storage

import (
	"context"
	"indetity-service/modules"
	"indetity-service/modules/item/entity"
)

func (sql *SqlStorage) GetPermission(ctx context.Context, cond map[string]interface{}) (*entity.Permission, error) {
	var data entity.Permission

	if err := sql.db.Where("name = ?", cond["name"]).First(&data).Error; err != nil {
		return nil, entity.PermissionNotFound
	}

	return &data, nil
}

func (sql *SqlStorage) GetAllPermission(ctx context.Context, paging *modules.Paging) ([]entity.Permission, error) {
	var data []entity.Permission

	paging.Process()

	if err := sql.db.Table(entity.Permission{}.TableName()).Count(&paging.TotalPage).Error; err != nil {
		return nil, entity.CantGetListPermission(err)
	}

	if err := sql.db.Offset((paging.Page - 1) * (paging.Size)).
		Limit(paging.Size).
		Find(&data).Error; err != nil {
		return nil, entity.CantGetListPermission(err)
	}

	return data, nil
}
