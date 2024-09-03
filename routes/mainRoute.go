package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
	"new-shout-golang/controllers"
	"new-shout-golang/middlewares"
	"new-shout-golang/repositories"
	"new-shout-golang/services"
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "HALO SAYANG")
	})

	// Swagger API Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	sellerRepo := repositories.NewSellerRepository(db)
	sellerService := services.NewSellerService(sellerRepo)
	sellerController := controllers.NewSellerController(sellerService)

	sellerRoutes := r.Group("/sellers")
	// Routes with auth middleware
	sellerRoutes.Use(middlewares.AuthMiddleware())
	{
		sellerRoutes.POST("/", sellerController.CreateSeller)
		sellerRoutes.GET("/", sellerController.GetAllSellers)
		sellerRoutes.GET("/:id", sellerController.GetSellerByID)
		sellerRoutes.PUT("/:id", sellerController.UpdateSeller)
		sellerRoutes.DELETE("/:id", sellerController.DeleteSeller)
	}
}
