package shipping_method_handle

import (
	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/shippingMethod"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"strconv"

	// "ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteShippingMethod(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz_shipping_method.NewDeleteShippingMethodBiz(store)

		if err := business.DeleteShippingMethodByID(c, id); err != nil {
			if err == models.ErrorNotFoundShippingMethod {
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

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
