package routes

import (
	"boodschappenlijst/controllers"
	"boodschappenlijst/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/login", controllers.Login)
        authRoutes.POST("/signup", controllers.Signup)
        authRoutes.GET("/home", controllers.Home)
        authRoutes.GET("/premium", controllers.Premium)
        authRoutes.GET("/logout", controllers.Logout)
    }
}

func DBRoutes(r *gin.Engine) {
    dbRoutes := r.Group("/db")
    {
        dbRoutes.POST("/products", middlewares.IsAuthorized(), middlewares.IsAdmin(), controllers.SaveProduct)
        dbRoutes.GET("/products", middlewares.IsAuthorized(), controllers.GetProducts)
        dbRoutes.PUT("/products", middlewares.IsAuthorized(), middlewares.IsAdmin(), controllers.UpdateProduct)
        dbRoutes.DELETE("/products", middlewares.IsAuthorized(), middlewares.IsAdmin(), controllers.DeleteProduct)
    }
}

func ViewRoutes(r *gin.Engine) {
    viewRoutes := r.Group("/view")
    {
        viewRoutes.GET("/products", middlewares.IsAuthorized(), nil)
        viewRoutes.GET("/login", controllers.ViewLogin)
    }

}