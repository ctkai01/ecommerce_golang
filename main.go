package main

import (
	"ecommerce_shop/configs"
	"ecommerce_shop/database"
	"ecommerce_shop/middlewares"

	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/modules/ecommerce/transport/gin/authHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/userHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/countriesHandle"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	configs.InitEnvConfigs()

	database.Connect(configs.EnvConfigs.DB)
	r := gin.Default()
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{	
			
			auth.POST("/register", auth_handler.CreateUser(database.DB))
			auth.POST("/login", auth_handler.LoginUser(database.DB))
		}
		authorize := api.Group("/", middlewares.AuthMiddleware)
		
		user := authorize.Group("/user") 
		{
			user.GET("/me", user_handler.CurrentUser(database.DB))
			user.PUT("/me", user_handler.UpdateUser(database.DB))
			user.PUT("/me/password", user_handler.ChangePassword(database.DB))
			user.POST("/me/logout", user_handler.LogOut(database.DB))
		}

		countries := authorize.Group("/countries") 
		{
			countries.GET("/", countries_handle.ListCountry(database.DB))
			countries.POST("/", middlewares.AdminMiddleware, countries_handle.CreateCountry(database.DB))
			countries.PUT("/:id", middlewares.AdminMiddleware, countries_handle.UpdateCountry(database.DB))
			countries.DELETE("/:id", middlewares.AdminMiddleware, countries_handle.DeleteCountry(database.DB))
			
			// countries.
			// user.PUT("/me", gin_ecommerce.UpdateUser(database.DB))
			// user.PUT("/me/password", gin_ecommerce.ChangePassword(database.DB))
			// user.POST("/me/logout", gin_ecommerce.LogOut(database.DB))
		}

	}

	bcrypt.GenerateFromPassword([]byte("Nam"), 10)
	r.Run(":8000")
}
