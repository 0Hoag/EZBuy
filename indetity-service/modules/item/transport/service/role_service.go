package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"indetity-service/dto"
	"indetity-service/modules"
	"indetity-service/modules/item/biz/role"
	"indetity-service/modules/item/entity"
	"indetity-service/modules/item/storage"
	"net/http"
)

func CreateRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data dto.RoleRequest

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		r := dto.ToRole(&data)

		store := storage.NewSQLStore(db)
		business := role.CreateRoleStorage(store)

		err := business.CreateRole(c.Request.Context(), r)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": r.Name,
		})
	}
}

func GetRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entity.Role

		id := c.Param("name")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := role.NewGetRole(store)

		data1, err := business.GetRoleBiz(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := dto.ToRoleResponse(*data1)

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}

func UpdateRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data dto.TodoRoleUpdate
		id := c.Param("name")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := role.NewUpdateRoleBiz(store)

		err := business.UpdateRoleBiz(c.Request.Context(), id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func DeleteRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("name")

		store := storage.NewSQLStore(db)
		business := role.NewDeleteRoleBiz(store)

		if err := business.DeleteRoleBiz(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func GetAllRole(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging modules.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := role.NewGetRole(store)

		data, err := business.GetAllRoleBiz(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"paging": paging,
			"result": data,
		})
	}
}
