package gin_ecommerce

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateUserBiz(store)

		if err := business.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(10))
	}
}