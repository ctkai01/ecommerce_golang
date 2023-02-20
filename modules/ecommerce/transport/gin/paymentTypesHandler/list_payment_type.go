package payment_type_handle

import (
	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/paymentType"
	"ecommerce_shop/modules/ecommerce/storage"

	// "ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListPaymentType(db *gorm.DB) func(*gin.Context) {
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

		business := biz_payment_type.ListPaymentTypeBiz(store)

		result, err := business.ListPaymentType(c, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
