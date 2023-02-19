package countries_handle

import (
	"ecommerce_shop/common"
	"ecommerce_shop/modules/ecommerce/biz/countries"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCountry(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.CreateCountry

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

		business := biz_countries.NewCreateCountryBiz(store)

		if err := business.CreateCountry(c, &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}


		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))
	}
}
