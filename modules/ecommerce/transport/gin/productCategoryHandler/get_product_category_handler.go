package product_category_handle

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/productCategory"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return 
		}

		store := storage.NewSQLStore(db)

		business := biz_product_category.NewGetProductCategoryBiz(store)

		data, err := business.GetProductCategory(c, id)

		if err != nil {
			if err == models.ErrorNotFoundProductCategory {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return 
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return 
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
} 