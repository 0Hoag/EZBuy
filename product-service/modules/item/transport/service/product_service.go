package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"product-service/dto"
	modules "product-service/modules/item"
	"product-service/modules/item/biz/products"
	"product-service/modules/item/storage"
)

func CreateProduct(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request dto.ProductRequest

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		product := request.ToProduct()
		if err := product.BeforeCreate(db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Create Product UUID it found!",
				"error":   err,
			})
			return
		}

		store := storage.NewSqlStore(db)
		business := products.CreateProductStorage(store)

		if err := business.CreateProduct(c.Request.Context(), product); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(product.Id))
	}
}

func GetProduct(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := storage.NewSqlStore(db)
		business := products.NewGetCategoryBiz(store)

		err, data := business.GetProductByIdBiz(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := dto.ToProductResponse(data)

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(response))
	}
}

func GetAllProduct(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging modules.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStore(db)
		business := products.NewGetCategoryBiz(store)

		err, data := business.GetAllProductBiz(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.NewSuccessResponse(data, paging))
	}
}

func DeleteProduct(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		store := storage.NewSqlStore(db)
		business := products.NewDeleteCategory(store)

		if err := business.DeleteCategoryBiz(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}

func UpdateProduct(db *mongo.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var updateRequest *dto.ProductUpdateRequest
		id := c.Param("id")

		if err := c.ShouldBind(&updateRequest); err != nil {
			c.JSON(http.StatusBadRequest, modules.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSqlStore(db)
		business := products.NewUpdateCategory(store)

		if err := business.UpdateProductBiz(c.Request.Context(), id, updateRequest); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, modules.SimpleSuccessResponse(true))
	}
}
