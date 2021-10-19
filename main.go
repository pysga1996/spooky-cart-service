package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pysga1996/spooky-cart-service/config"
	"github.com/pysga1996/spooky-cart-service/controller"
	"github.com/pysga1996/spooky-cart-service/middleware"
)

func main() {
	db := config.ConnectDatabase() //Kết nối database
	controller.DB = db
	// close database
	defer func(db *sql.DB) {
		err := db.Close()
		middleware.CheckErrorShutdown(err)
	}(db)
	router := gin.Default()
	router.Use(middleware.HandleError, middleware.HandleToken)

	RegisterRoutes(router)

	// Start serving the application
	err2 := router.Run()
	middleware.CheckErrorShutdown(err2)

}
