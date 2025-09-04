package main

import (
	"os"
	"time"

	"github.com/Shashank100909/STUDENTS-API/configs"
	dao "github.com/Shashank100909/STUDENTS-API/daos"
	handlers "github.com/Shashank100909/STUDENTS-API/handler"
	middleware "github.com/Shashank100909/STUDENTS-API/middlewere"
	"github.com/Shashank100909/STUDENTS-API/models"
	services "github.com/Shashank100909/STUDENTS-API/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	configs.DB.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Products{},
		&models.Address{},
	)

	configs.DB.AutoMigrate(
		&models.Cart{},
	)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://react-projects-beta-hazel.vercel.app",
		},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userDAO := dao.NewUserDAO()
	userService := services.NewUserService(userDAO)
	userHandler := handlers.NewUserHandler(userService)

	productDAO := dao.NewProductDAO()
	productService := services.NewProductService(productDAO)
	productHandler := handlers.NewProductHandler(productService, productDAO)

	r.POST("/register", userHandler.Register)
	r.POST("/student-register", userHandler.CreateStudent)
	r.GET("/student-register", userHandler.GetStudent)
	r.POST("/login", userHandler.Login)
	r.DELETE("student-delete/:id", userHandler.DeleteStudent)

	auth := r.Group("/", middleware.AuthMiddleware())

	auth.POST("/product", productHandler.AddProduct)
	auth.GET("/product", productHandler.GetProducts)
	auth.DELETE("/product/:id", productHandler.DeleteProduct)

	auth.POST("/cart", productHandler.AddProductToCart)
	auth.GET("/cart/:user_id", productHandler.GetCartItems)
	auth.DELETE("cart/:product_id", productHandler.DeleteProductFromCart)

	auth.POST("address", productHandler.AddAddress)

	// auth.PUT("/cart/:user_id", productHandler.UpdateCartItems)

	// r.Run(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
