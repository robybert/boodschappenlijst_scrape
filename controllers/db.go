package controllers

import (
	"boodschappenlijst/db"
	"boodschappenlijst/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetProducts(c *gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		// return err
	}

	retProducts := db.Database.GetProducts(product)
	c.JSON(http.StatusOK, retProducts)
}

func SaveProduct(c *gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {

	}

	c.JSON(http.StatusOK, db.Database.SaveProduct(product))
}

func UpdateProduct(c *gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {

	}
	updatedProduct := db.Database.UpdateProduct(product)

	c.JSON(http.StatusOK, updatedProduct)
}
func DeleteProduct(c *gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {

	}

	if !db.Database.DeleteProduct(product) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "product deleted."})
}


