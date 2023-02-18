package database

import (
	// "ecommerce_shop/modules/ecommerce/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect(connectString string) {
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// db.SetupJoinTable(&models.User{}, "Addresses", &models.UserAddress{})
	// db.SetupJoinTable(&models.User{}, "PaymentTypes", &models.UserPaymentMethod{})
	// db.AutoMigrate(
	// 	&models.Country{}, 
	// 	&models.Address{}, 
	// 	&models.User{}, 
	// 	&models.PaymentType{},
	// 	&models.ShoppingCart{}, 
	// )

	DB = db
}
