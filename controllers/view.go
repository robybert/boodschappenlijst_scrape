package controllers

import (
	"boodschappenlijst/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewProducts(c *gin.Context){
	products := db.Database.GetAllProducts()

	data := gin.H{	"title"		: 	"Product Database",
					"products"	:	products,
	}
	c.HTML(http.StatusOK, "index.html", data)
}

func ViewLogin(c *gin.Context) {
	
}