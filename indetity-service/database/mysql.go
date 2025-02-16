package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"indetity-service/modules/item/entity"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDatabase() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := "root:601748@tcp(127.0.0.1:3307)/identity-service?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}

		err = db.AutoMigrate(
			&entity.User{},
			&entity.Role{},
			&entity.Permission{},
			&entity.UserRoles{},
			&entity.RolePermissions{},
		)

		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	})

	return db, err
}

func getDB() *gorm.DB {
	return db
}
