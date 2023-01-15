package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/splash-inventory-service/internal/config"
	middleware2 "github.com/thanh-vt/splash-inventory-service/internal/middleware"
)

func main() {
	config.ConnectDatabase() // connect database

	router := gin.Default()
	router.Use(middleware2.HandleError, middleware2.HandleToken)

	config.RegisterRoutes(router)
	// Start serving the application
	err2 := router.Run()
	middleware2.CheckErrorShutdown(err2)

}
