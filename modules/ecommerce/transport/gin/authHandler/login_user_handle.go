package auth_handler

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/auth"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validateErr := validate.Struct(data)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz_auth.NewLoginUserBiz(store)

		
		if err := business.LoginUser(c.Request.Context(), &data); err != nil {
			if err == common.ErrorInvalidCredentials {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(
			data,
		))
	}
}