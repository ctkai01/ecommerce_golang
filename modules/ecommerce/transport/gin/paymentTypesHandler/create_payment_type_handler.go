package payment_type_handle

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/paymentType"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePaymentType(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.CreatePaymentType

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

		business := biz_payment_type.NewCreatePaymentTypeBiz(store)

		if err := business.CreatePaymentType(c, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}


		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))
	}
}
