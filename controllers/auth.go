package controllers

import (
	"boodschappenlijst/db"
	"boodschappenlijst/models"
	"fmt"
	"time"

	"boodschappenlijst/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// The string "my_secret_key" is just an example and should be replaced with a secret key of sufficient length and complexity in a real-world scenario.
var jwtKey = []byte("my_secret_key")

func Login(c *gin.Context) {

    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	existingUser := db.Database.GetUserByMail(user)

    if existingUser.ID == 0 {
        c.JSON(400, gin.H{"error": "user does not exist"})
        return
    }

    errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

    if !errHash {
        c.JSON(400, gin.H{"error": "invalid password"})
        return
    }

    expirationTime := time.Now().Add(5 * time.Minute)

    claims := &models.Claims{
        Role: existingUser.Role,
        StandardClaims: jwt.StandardClaims{
            Subject:   existingUser.Email,
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(jwtKey)

    if err != nil {
        c.JSON(500, gin.H{"error": "could not generate token"})
        return
    }

    c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
    c.JSON(200, gin.H{"success": "user logged in"})
}

func Signup(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"errorh": err.Error()})
        return
    }

    existingUser := db.Database.GetUserByMail(user)

    if existingUser.ID != 0 {
        c.JSON(400, gin.H{"error": "user already exists"})
        return
    }

    var errHash error
    user.Password, errHash = utils.GenerateHashPassword(user.Password)

    if errHash != nil {
        c.JSON(500, gin.H{"error": "could not generate password hash"})
        return
    }

	if user.Role != "user" {
		c.JSON(401, gin.H{"error": "not authorized to create admin"})
		return
	}

	if !db.Database.SaveUser(user) {
		//TODO:  error handling
	}

    c.JSON(200, gin.H{"success": "user created"})
}

func Home(c *gin.Context) {

    cookie, err := c.Cookie("token")

    if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    claims, err := utils.ParseToken(cookie)

    if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

	fmt.Printf("%s\n", claims.Role)

    if claims.Role != "user" && claims.Role != "admin" {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func Premium(c *gin.Context) {

    cookie, err := c.Cookie("token")

    if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    claims, err := utils.ParseToken(cookie)

    if err != nil {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    if claims.Role != "admin" {
        c.JSON(401, gin.H{"error": "unauthorized"})
        return
    }

    c.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
}

func Logout(c *gin.Context) {
    c.SetCookie("token", "", -1, "/", "localhost", false, true)
    c.JSON(200, gin.H{"success": "user logged out"})
}

// func Products(c *gin.Context) {
// 	middlewares.IsAuthorized()

// 	var product models.Product

// 	if c.Request.Method == http.MethodPost {
// 		// if err:= c.ShouldBindJSON(&product); err != nil{
// 		// 	fmt.Printf("err: %v\n", err)
// 		// }

// 		// fmt.Printf("product: %s\n", product.ProductName)
// 		// var retProd models.Product

// 		// models.DB.Where("product_name = ?", product.ProductName).First(&retProd)

// 		// retJSON, _ := json.Marshal(retProd)


// 		// fmt.Printf("retProduct: %s\n", retProd.ProductName)



// 		// c.JSON(200, gin.H{retJSON})
// 	} else if c.Request.Method == http.MethodGet {
// 		if err:= c.ShouldBindJSON(&product); err != nil{
// 			fmt.Printf("err: %v\n", err)
// 		}

// 		var retProds []models.Product

// 		models.DB.Where("product_name = ?", product.ProductName).Find(&retProds)

// 		if json, err := json.Marshal(retProds); err != nil {
// 			fmt.Printf("err: %v\n", err)
// 		}

		
// 		c.JSON(http.StatusOK, gin.H{&json})

// 	}

// 	if err:= c.ShouldBindJSON(&product); err != nil{
// 		fmt.Printf("err: %v\n", err)
// 	}

	



// }