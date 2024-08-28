package main

import (
	"new-shout-golang/api"
)

// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	api.App.Run()
}
