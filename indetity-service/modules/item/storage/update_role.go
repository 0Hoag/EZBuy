package storage

import (
	"context"
	"indetity-service/dto"
	"indetity-service/modules/item/entity"
)

func (s *SqlStorage) UpdateRole(ctx context.Context, cond map[string]interface{}, dataUpdate *dto.TodoRoleUpdate) error {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var role entity.Role
	if err := tx.Preload("Permission").Where("name = ?", cond["name"]).First(&role).Error; err != nil {
		tx.Rollback()
		return entity.RoleNotFound
	}

	if dataUpdate.Name != nil {
		role.Name = *dataUpdate.Name
	}

	if dataUpdate.Description != nil {
		role.Description = *dataUpdate.Description
	}

	if dataUpdate.Permission != nil {
		var permissions []entity.Permission
		if err := tx.Where("name IN (?)", *dataUpdate.Permission).Find(&permissions).Error; err != nil {
			tx.Rollback()
			return entity.CantGetListPermission(err)
		}

		if err := tx.Model(&role).Association("Permission").Clear(); err != nil {
			tx.Rollback()
			return entity.CantGetListPermission(err)
		}

		if err := tx.Model(&role).Association("Permission").Append(permissions); err != nil {
			tx.Rollback()
			return entity.CantGetListPermission(err)
		}
	}

	if err := tx.Save(&role).Error; err != nil {
		tx.Rollback()
		return entity.CantUpdateRole(err)
	}

	return tx.Commit().Error
}
