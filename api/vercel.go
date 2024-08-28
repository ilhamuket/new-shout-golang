package api

import (
	"log"
	"net/http"
	"new-shout-golang/config"
	"new-shout-golang/docs"
	"new-shout-golang/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func init() {
	App = gin.New()

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "E-Commerce REST API"
	docs.SwaggerInfo.Description = "This is REST API E-Commerce."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	if docs.SwaggerInfo.Host == "" {
		docs.SwaggerInfo.Host = "localhost:8080"
	}
	if environment == "development" {
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	db := config.ConnectDataBase()
	routes.SetupRouter(db, App)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
