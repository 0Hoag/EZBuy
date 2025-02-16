package config

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"indetity-service/modules/item/entity"
	"log"
)

const (
	ADMIN_USER_NAME = "admin"
	ADMIN_PASSWORD  = "admin"
)

func InitializeAdminUser(db *gorm.DB) error {
	var existingUserRole, existingAdminRole entity.Role
	userRoleExists := db.Where("name = ?", "USER").First(&existingUserRole).Error == nil
	adminRoleExists := db.Where("name = ?", "ADMIN").First(&existingAdminRole).Error == nil

	if !userRoleExists {
		userRole := entity.Role{
			Name:        "USER",
			Description: "User role",
		}
		if err := db.Create(&userRole).Error; err != nil {
			return fmt.Errorf("failed to create USER role: %v", err)
		}
	}

	if !adminRoleExists {
		adminRole := entity.Role{
			Name:        "ADMIN",
			Description: "Admin role",
		}
		if err := db.Create(&adminRole).Error; err != nil {
			return fmt.Errorf("failed to create ADMIN role: %v", err)
		}
	}

	var existingUser entity.User
	userExists := db.Where("username = ?", ADMIN_USER_NAME).First(&existingUser).Error == nil

	if !userExists {
		var adminRole entity.Role
		if err := db.Where("name = ?", "ADMIN").First(&adminRole).Error; err != nil {
			return fmt.Errorf("failed to retrieve ADMIN role: %v", err)
		}

		adminUser := entity.User{
			UserId:   uuid.New().String(),
			Username: ADMIN_USER_NAME,
			Password: ADMIN_PASSWORD,
			Role:     []entity.Role{adminRole},
		}

		fmt.Println("user: {}", adminUser)
		fmt.Println("role: {}", adminUser.Role)

		//if err := db.Table(models.UserRoles{}.TableName()).Create(&userRole).Error; err != nil {
		//	return fmt.Errorf("failed to create user_roles: %v", err)
		//}

		if err := db.Create(&adminUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %v", err)
		}

		log.Println("Admin user has been created with default password: admin, please change it")
	} else {
		log.Println("Admin user already exists")
	}

	return nil
}
