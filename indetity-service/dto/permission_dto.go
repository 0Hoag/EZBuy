package dto

import (
	"indetity-service/modules/item/entity"
)

type PermissionRequest struct {
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

type PermissionResponse struct {
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

type TodoPermissionUpdate struct {
	Description *string `json:"description" gorm:"description"`
}

func (TodoPermissionUpdate) TableName() string {
	return "permission"
}

func ToPermission(req *PermissionRequest) *entity.Permission {
	return &entity.Permission{
		Name:        req.Name,
		Description: req.Description,
	}
}

func ToPermissionResponse(p *entity.Permission) *PermissionResponse {
	return &PermissionResponse{
		Name:        p.Name,
		Description: p.Description,
	}
}
