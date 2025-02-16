package main

import (
	"github.com/gin-gonic/gin"
	"indetity-service/config"
	"indetity-service/database"
	"indetity-service/modules/item/transport/service"
	"log"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	err = config.InitializeAdminUser(db)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()

	scope := r.Group("/v1/api") // scope
	{
		user := scope.Group("/user")
		{
			user.POST("", service.CreateUser(db))
			user.GET("/getAll", service.GetAllUser(db))
			user.GET("/:user_id", service.GetUser(db))
			user.GET("/getUserName/:username", service.GetUserByUserName(db))
			user.PATCH("/:user_id", service.UpdateUser(db))
			user.DELETE("/:user_id", service.DeleteUser(db))
		}

		role := scope.Group("/role")
		{
			role.POST("", service.CreateRole(db))
			role.GET("/getAll", service.GetAllRole(db))
			role.GET("/:name", service.GetRole(db))
			role.PUT("/:name", service.UpdateRole(db))
			role.DELETE("/:name", service.DeleteRole(db))
		}

		permission := scope.Group("/permission")
		{
			permission.POST("", service.CreatePermission(db))
			permission.GET("/getAll", service.GetAllPermission(db))
			permission.GET("/:name", service.GetPermission(db))
			permission.PATCH("/:name", service.UpdatePermission(db))
			permission.DELETE("/:name", service.DeletePermission(db))
		}

		authentication := scope.Group("/auth")
		{
			authentication.POST("/token", service.Login(db))
			authentication.POST("/logout", service.Logout(db))
			authentication.POST("/refresh", service.RefreshToken(db))
		}
	}

	r.Run(":8081")
}
