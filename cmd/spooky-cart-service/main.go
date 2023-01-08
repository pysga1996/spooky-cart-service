package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/spooky-cart-service/internal/config"
	"github.com/thanh-vt/spooky-cart-service/internal/controller"
	middleware2 "github.com/thanh-vt/spooky-cart-service/internal/middleware"
)

func main() {
	db := config.ConnectDatabase() // connect database
	controller.DB = db
	// close database
	defer func(db *sql.DB) {
		err := db.Close()
		middleware2.CheckErrorShutdown(err)
	}(db)
	router := gin.Default()
	router.Use(middleware2.HandleError, middleware2.HandleToken)

	config.RegisterRoutes(router)
	// Start serving the application
	err2 := router.Run()
	middleware2.CheckErrorShutdown(err2)

}
