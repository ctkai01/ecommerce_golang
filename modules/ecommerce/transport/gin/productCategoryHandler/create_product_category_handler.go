package product_category_handle

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/productCategory"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProductCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.CreateProductCategory

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validateErr := Validate.Struct(data)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz_product_category.NewCreateProductCategory(store)

		if err := business.CreateProductCategory(c, &data); err != nil {
			if err == models.ErrorNotFoundProductCategory {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return 
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.ID))
	}
}
