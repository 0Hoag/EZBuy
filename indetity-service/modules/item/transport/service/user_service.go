package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"indetity-service/dto"
	"indetity-service/modules"
	"indetity-service/modules/item/biz/user"
	"indetity-service/modules/item/entity"
	"indetity-service/modules/item/storage"
	"net/http"
)

func CreateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.UserRequest

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		var us = request.ToUser()
		if err := us.BeforeCreate(db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Create User UUID it found!",
				"error":   err,
			})
			return
		}

		var role []entity.Role
		if err := db.Where("name = ?", "USER").First(&role).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": entity.UserNotFound,
				"error":   err,
			})
			return
		}

		us.Role = role

		store := storage.NewSQLStore(db)
		business := user.NewCreateUserBiz(store)

		if err := business.CreateNewBiz(c.Request.Context(), us); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(us.UserId))
	}
}

func GetUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("user_id")

		store := storage.NewSQLStore(db)
		business := user.NewGetUserBiz(store)

		data, err := business.GetUserByIdBiz(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := dto.ToUserResponse(data)

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(response))
	}
}

func GetUserByUserName(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.Param("username")

		store := storage.NewSQLStore(db)
		business := user.NewGetUserBiz(store)

		data, err := business.GetUserByUserNameBiz(c.Request.Context(), username)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := dto.ToUserResponse(data)

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(response))
	}
}

func UpdateUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data dto.TodoUpdateRequest

		id := c.Param("user_id")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := user.NewUpdateUserBiz(store)

		if err := business.UpdateUserBiz(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}

func DeleteUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("user_id")

		store := storage.NewSQLStore(db)
		business := user.NewDeleteUserBiz(store)

		if err := business.DeleteUserBiz(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}

func GetAllUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging modules.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := user.NewGetUserBiz(store)

		data, err := business.GetAllUser(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.NewSuccessResponse(data, paging))
	}
}
