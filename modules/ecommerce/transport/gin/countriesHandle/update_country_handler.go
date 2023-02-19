package countries_handle

import (
	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/common"
	biz_countries "ecommerce_shop/modules/ecommerce/biz/countries"
	"ecommerce_shop/modules/ecommerce/models"
	"ecommerce_shop/modules/ecommerce/storage"
	"strconv"

	// "ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCountry(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.UpdateCountry

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&data); err != nil {
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

		store := storage.NewSQLStore(db)

		business := biz_countries.NewUpdateCountryBiz(store)

		if err := business.UpdateCountryByID(c, id, &data); err != nil {
			if err == models.ErrorNotFoundCountry {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(true))
	}
}
