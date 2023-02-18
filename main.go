package main

import (
	"ecommerce_shop/configs"
	"ecommerce_shop/database"
	
	"ecommerce_shop/modules/ecommerce/transport/gin"
	
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitEnvConfigs()
	
	database.Connect(configs.EnvConfigs.DB)
	r := gin.Default()
	api := r.Group("/api") 
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register",  gin_ecommerce.CreateUser(database.DB))
		}
		
	}
	
	
	bcrypt.GenerateFromPassword([]byte("Nam"), 10)
	r.Run(":8000")
}
