package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"indetity-service/dto"
	"indetity-service/modules"
	"indetity-service/modules/item/biz/permission"
	"indetity-service/modules/item/entity"
	"indetity-service/modules/item/storage"
	"net/http"
)

func CreatePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data dto.PermissionRequest

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		r := dto.ToPermission(&data)

		store := storage.NewSQLStore(db)
		business := permission.CreatePermissionStorage(store)

		if err := business.CreatePermission(c.Request.Context(), r); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		c.JSON(http.StatusOK, gin.H{
			"response": r.Name,
		})
	}
}

func GetPermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data entity.Permission

		id := c.Param("name")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := permission.NewGetPermission(store)

		result, err := business.GetPermissionBiz(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"response": result,
		})
	}
}

func UpdatePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data dto.TodoPermissionUpdate

		id := c.Param("name")

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := permission.NewUpdatePermission(store)

		if err := business.UpdatePermissionBiz(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}

func DeletePermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("name")

		store := storage.NewSQLStore(db)
		business := permission.NewDeletePermission(store)

		if err := business.DeletePermissionBiz(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func GetAllPermission(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging modules.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := permission.NewGetPermission(store)

		data, err := business.GetAllPermissionBiz(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"paging": paging,
			"result": data,
		})
	}
}
