package middlewares

import (
	"ecommerce_shop/modules/ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	dataUserAuth, isExist := c.Get("user")

	if !isExist {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	dataUser := dataUserAuth.(models.User)

	if dataUser.IsAdmin {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}
