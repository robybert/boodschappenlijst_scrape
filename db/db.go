package db

import (
	"boodschappenlijst/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB interface {
	GetProducts(product models.Product) []models.Product
	SaveProduct(product models.Product) models.Product
	UpdateProduct(product models.Product) models.Product
	DeleteProduct(product models.Product) bool
	GetUserByMail(user models.User) models.User
	SaveUser(user models.User) bool
	UpdateUser(user models.User) bool
	DeleteUser(user models.User) bool
	CloseDB()
}

type database struct {
	DB *gorm.DB
}

 
func NewDatabase() DB {
	db, err := gorm.Open(sqlite.Open("db/DB.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Price{}); err != nil {
        panic(err)
    }

    fmt.Println("Migrated database")

	return &database {
		DB: db,
	}

}

var Database DB = NewDatabase()

func (db *database) GetProducts(product models.Product) []models.Product {
	var retProds []models.Product

	db.DB.Where("product_name = ?", product.ProductName).Find(&retProds)
	return retProds
}

func (db *database) SaveProduct(product models.Product) models.Product {
	db.DB.Create(&product)
	if db.DB.Error != nil {
		
	}
	return product
}

func (db *database) UpdateProduct(product models.Product) models.Product {
	db.DB.Model(&product).Updates(&product)
	if db.DB.Error != nil {
		
	}
	var updatedProduct models.Product
	db.DB.First(&updatedProduct, product.ID)
	return updatedProduct
}
func (db *database) DeleteProduct(product models.Product) bool {
	db.DB.Delete(&product)
	if db.DB.Error != nil {

	}
	return true
}
func (db *database) CloseDB() {
	closeDB, err := db.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer closeDB.Close()
	
}

//TODO: error checking

func (db *database) GetUserByMail(user models.User) models.User{
	var existingUser models.User
	db.DB.Where("email = ?", user.Email).First(&existingUser)
	return existingUser
}

func (db *database) SaveUser(user models.User) bool{
	db.DB.Create(&user)
	if db.DB.Error != nil {

	}
	return true
}
func (db *database) UpdateUser(user models.User) bool{
	db.DB.Save(&user)
	if db.DB.Error != nil {

	}
	return true
}
func (db *database) DeleteUser(user models.User) bool{
	db.DB.Delete(&user)
	if db.DB.Error != nil {

	}
	return true
}