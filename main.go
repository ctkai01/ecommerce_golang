package main

import (
	"ecommerce_shop/configs"
	"ecommerce_shop/database"
	"ecommerce_shop/middlewares"

	// "ecommerce_shop/modules/ecommerce/models"

	"ecommerce_shop/modules/ecommerce/transport/gin/authHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/userHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/countriesHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/paymentTypesHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/shippingMethodHandler"
	"ecommerce_shop/modules/ecommerce/transport/gin/productCategoryHandler"

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
		}

		paymentTypes := authorize.Group("/payment_types") 
		{
			paymentTypes.GET("/", payment_type_handle.ListPaymentType(database.DB))
			
			paymentTypes.POST("/", middlewares.AdminMiddleware, payment_type_handle.CreatePaymentType(database.DB))
			paymentTypes.PUT("/:id", middlewares.AdminMiddleware, payment_type_handle.UpdatePaymentType(database.DB))
			paymentTypes.DELETE("/:id", middlewares.AdminMiddleware, payment_type_handle.DeletePaymentType(database.DB))
		}

		shoppingMethod := authorize.Group("/shopping_methods") 
		{
			shoppingMethod.GET("/", shipping_method_handle.ListShippingMethod(database.DB))
			shoppingMethod.POST("/", middlewares.AdminMiddleware, shipping_method_handle.CreateShippingMethod(database.DB))
			shoppingMethod.PUT("/:id", middlewares.AdminMiddleware, shipping_method_handle.UpdateShippingMethod(database.DB))
			shoppingMethod.DELETE("/:id", middlewares.AdminMiddleware, shipping_method_handle.DeleteShippingMethod(database.DB))
		}

		productCategory := authorize.Group("/product_categories") 
		{	
			productCategory.GET("/", product_category_handle.ListProductCategory(database.DB))
			productCategory.GET("/:id", product_category_handle.GetProductCategory(database.DB))
			productCategory.POST("/", middlewares.AdminMiddleware, product_category_handle.CreateProductCategory(database.DB))
			productCategory.PUT("/:id", middlewares.AdminMiddleware, product_category_handle.UpdateProductCategory(database.DB))
			productCategory.DELETE("/:id", middlewares.AdminMiddleware, product_category_handle.DeleteProductCategory(database.DB))
		}

	}

	bcrypt.GenerateFromPassword([]byte("Nam"), 10)
	r.Run(":8080")
}
