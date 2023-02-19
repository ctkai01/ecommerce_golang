package auth_handler

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/auth"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data models.UserCreation

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

		business := biz_auth.NewCreateUserBiz(store)

		if err := business.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.Id))
	}
}
