package user_handler

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/user"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangePassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserUpdatePassword

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

		dataUserAuth, isExist := c.Get("user")

		if !isExist {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		dataUser := dataUserAuth.(models.User)

		store := storage.NewSQLStore(db)
		
		business := biz_user.NewChangePasswordBiz(store)

		if err := business.ChangePassword(c.Request.Context(), dataUser.ID, &data); err != nil {
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

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}

}