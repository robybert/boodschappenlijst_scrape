package routes

import (
	"boodschappenlijst/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/login", controllers.Login)
    r.POST("/signup", controllers.Signup)
    r.GET("/home", controllers.Home)
    r.GET("/premium", controllers.Premium)
    r.GET("/logout", controllers.Logout)
}

func DBRoutes(r *gin.Engine) {
    r.POST("/products", controllers.SaveProduct)
    r.GET("/products", controllers.GetProducts)
    r.PUT("/products", controllers.UpdateProduct)
    r.DELETE("/products", controllers.DeleteProduct)
}