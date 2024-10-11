package main

import (
	"boodschappenlijst/routes"

	"github.com/gin-gonic/gin"
)




func main() {
	r := gin.Default()

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// config := models.Config{
	// 	Host:     os.Getenv("DB_HOST"),
    //     Port:     os.Getenv("DB_PORT"),
    //     User:     os.Getenv("DB_USER"),
    //     Password: os.Getenv("DB_PASSWORD"),
    //     DBName:   os.Getenv("DB_NAME"),
    //     SSLMode:  os.Getenv("DB_SSLMODE"),
	// }

	routes.AuthRoutes(r)
	routes.DBRoutes(r)
	routes.ViewRoutes(r)

	r.Static("/static", "./templates/static")

	r.LoadHTMLGlob("templates/*.html")

	r.Run(":8080")
}