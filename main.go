package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type User struct {
	ID          int    `json:"id"`
	Fullname    string `json:"fullname"`
	Email    string `json:"email"`
	Password    string `json:"password"`
	// PhoneNumber string `json:"phone_number"`
}

var db *gorm.DB

func initDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/be_ucommerce" // Ganti dengan DSN MySQL Anda
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal membuka koneksi ke database: %v", err)
	}

	// Migrasi skema database
	db.AutoMigrate(&Product{}, &User{})
}
// product
func getProducts(c *gin.Context) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data produk"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func getOneProducts(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := db.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk"})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui produk"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

func handlerUser(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pengguna"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func handlerRegister(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil"})
}	

func handlerLogin(c *gin.Context) {
	var loginUser User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	var user User
	if err := db.Where("email = ?", loginUser.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	if loginUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "user": user})
}

func main() {
	initDB()

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

	// Product
	r.GET("/api/products", getProducts)
	r.GET("/api/products/:id", getOneProducts)
	r.POST("/api/products", createProduct)
	r.PUT("/api/products/:id", updateProduct)
	r.DELETE("/api/products/:id", deleteProduct)



	r.POST("/api/register", handlerRegister)
	r.POST("/api/login", handlerLogin)

	log.Println("Server berjalan di port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
