package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int   `json:"id"`
	ProductName string `json:"product_name"`
	Price       int   `json:"price"`
	Description string `json:"description"`
}

type Users struct {
	ID          int   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

var products = []Product{
	{ID: 1, ProductName: "Product 1", Price: 10000, Description: "Description 1"},
	{ID: 2, ProductName: "Product 2", Price: 20000, Description: "Description 2"},
	{ID: 3, ProductName: "Product 2", Price: 20000, Description: "Description 2"},
	{ID: 4, ProductName: "Product 2", Price: 20000, Description: "Description 2"},
	{ID: 5, ProductName: "Product 2", Price: 20000, Description: "Description 2"},
}

var users = []Users{
	{ID: 1, Username: "User1", Password: "user111", PhoneNumber: "08929393793"},
	{ID: 2, Username: "User2", Password: "user222", PhoneNumber: "08400750575"},
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

	r.GET("/api/users", handlerUser)
	r.GET("/api/products", handlerProduct)
	

	log.Println("Server berjalan di port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
