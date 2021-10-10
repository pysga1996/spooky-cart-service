package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pysga1996/spooky-cart-service/config"
	"github.com/pysga1996/spooky-cart-service/controller"
	"github.com/pysga1996/spooky-cart-service/error"
)

func main() {
	db := config.ConnectDatabase() //Kết nối database
	controller.DB = db
	// close database
	defer func(db *sql.DB) {
		err := db.Close()
		error.CheckErrorShutdown(err)
	}(db)
	router := gin.Default()

	router.Use(error.Handle, gin.Logger())
	RegisterRoutes(router)
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())

	// Start serving the application
	err2 := router.Run()
	error.CheckErrorShutdown(err2)

}
