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


func CurrentUser(db *gorm.DB) func(*gin.Context) {
	return func (c *gin.Context)  {
		var data models.UserAuth

		dataUserAuth, isExist := c.Get("user")

		if !isExist {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		dataUser := dataUserAuth.(models.User)

		store := storage.NewSQLStore(db)

		business := biz_user.NewCurrentUserBiz(store)
		

		if err := business.CurrentUser(c.Request.Context(), dataUser.ID, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}