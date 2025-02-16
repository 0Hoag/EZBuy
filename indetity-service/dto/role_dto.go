package dto

import (
	"indetity-service/modules/item/entity"
)

type RoleRequest struct {
	Name        string   `json:"name" gorm:"column:name;primaryKey"`
	Description string   `json:"description" gorm:"column:description"`
	Permission  []string `json:"permission" gorm:"many2many:role_permissions"`
}

type RoleResponse struct {
	Name        string              `json:"name" gorm:"column:name;primaryKey"`
	Description string              `json:"description" gorm:"column:description"`
	Permission  []entity.Permission `json:"permission" gorm:"many2many:role_permissions"`
}

type TodoRoleUpdate struct {
	Name        *string   `json:"name" gorm:"column:name;primaryKey"`
	Description *string   `json:"description" gorm:"column:description"`
	Permission  *[]string `json:"permission" gorm:"many2many:role_permissions"`
}

func ToRole(req *RoleRequest) *entity.Role { // mapstruct
	permissions := make([]entity.Permission, len(req.Permission))
	for i, p := range req.Permission {
		permissions[i] = entity.Permission{Name: p}
	}

	return &entity.Role{
		Name:        req.Name,
		Description: req.Description,
		Permission:  permissions,
	}
}

func ToRoleResponse(role entity.Role) *RoleResponse { // mapstruct
	return &RoleResponse{
		Name:        role.Name,
		Description: role.Description,
		Permission:  role.Permission,
	}
}
