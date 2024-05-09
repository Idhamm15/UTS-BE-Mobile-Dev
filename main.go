package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          uint   `json:"id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
}

type Users struct {
	ID          uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var products = []Product{
	{ID: 1, ProductName: "Product 1", Price: 100, Description: "Description 1"},
	{ID: 2, ProductName: "Product 2", Price: 200, Description: "Description 2"},
}

var users = []Users{
	{ID: 1, Username: "User1", Password: "user111"},
	{ID: 2, Username: "User2", Password: "user222"},
}

func handlerProduct(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func handlerUser(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func main() {
	r := gin.Default()

	// Middleware untuk menambahkan header CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/api/user", handlerProduct)
	r.GET("/api/products", handlerProduct)
	

	log.Println("Server berjalan di port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
