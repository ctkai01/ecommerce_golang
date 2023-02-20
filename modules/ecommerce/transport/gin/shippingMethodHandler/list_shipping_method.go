package shipping_method_handle

import (
	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/shippingMethod"
	"ecommerce_shop/modules/ecommerce/storage"

	// "ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListShippingMethod(db *gorm.DB) func(*gin.Context) {
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

		business := biz_shipping_method.ListShippingMethodBiz(store)

		result, err := business.ListShippingMethod(c, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
