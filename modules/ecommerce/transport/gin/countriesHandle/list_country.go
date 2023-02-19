package countries_handle

import (
	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/common"
	biz_countries "ecommerce_shop/modules/ecommerce/biz/countries"
	"ecommerce_shop/modules/ecommerce/storage"

	// "ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCountry(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		// var data models.CreateCountry
		var paging common.Paging


		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()
		store := storage.NewSQLStore(db)

		business := biz_countries.ListCountryBiz(store)

		result, err := business.ListCountry(c, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, common.NewSuccessResponse(result, paging, nil))
	}
}
