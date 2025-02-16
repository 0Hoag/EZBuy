package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"product-service/dto"
	modules "product-service/modules/item"
	"product-service/modules/item/biz/categories"
	"product-service/modules/item/entity"
	"product-service/modules/item/storage"
)

func CreateCategories(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.CategoryRequest

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		category := request.ToCategory()
		if err := category.BeforeCreate(db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Create User Categories it found!",
				"error":   err,
			})
			return
		}

		store := storage.NewSqlStore(db)
		business := categories.NewCreateCategory(store)

		if err := business.CreateCategoryBiz(c.Request.Context(), category); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(category.Id))
	}
}

func GetCategories(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := storage.NewSqlStore(db)
		business := categories.NewGetCategoryBiz(store)

		err, data := business.GetCategoryByIdBiz(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := dto.ToCategoryResponse(data)

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(response))
	}
}

func GetAllCategories(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging modules.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStore(db)
		business := categories.NewGetCategoryBiz(store)

		err, data := business.GetAllCategoryBiz(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.NewSuccessResponse(data, paging))
	}
}

func DeleteCategories(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := storage.NewSqlStore(db)
		business := categories.NewDeleteCategory(store)

		if err := business.DeleteCategoryBiz(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}

func UpdateCategories(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var updateRequest *dto.CategoryUpdateRequest
		id := c.Param("id")

		if err := c.ShouldBind(&updateRequest); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStore(db)
		business := categories.NewUpdateCategory(store)

		if err := business.UpdateCategoryBiz(c.Request.Context(), id, updateRequest); err != nil {
			c.JSON(http.StatusBadRequest, entity.CantUpdateCategory(err))
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}
