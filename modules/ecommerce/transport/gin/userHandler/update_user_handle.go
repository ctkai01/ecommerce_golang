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

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UserUpdate
	
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := Validate.Struct(data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		userAuth, isExist := c.Get("user")

		user := userAuth.(models.User)
		idUser := user.ID

		if !isExist {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		store := storage.NewSQLStore(db)
		business := biz_user.NewUpdateCurrentBiz(store)

		if err := business.UpdateCurrentUser(c, idUser, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
