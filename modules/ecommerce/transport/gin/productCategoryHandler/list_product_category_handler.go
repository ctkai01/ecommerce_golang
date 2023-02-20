package product_category_handle

import (

	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/productCategory"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProductCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()
		store := storage.NewSQLStore(db)

		business := biz_product_category.NewListProductCategoryBiz(store)

		result, err := business.ListProductCategory(c, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
