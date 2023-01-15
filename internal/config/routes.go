package config

import (
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/splash-inventory-service/internal/controller"
	"github.com/thanh-vt/splash-inventory-service/internal/middleware"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	// Set the router as the default one provided by Gin

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/css", "web/css")
	router.Static("/js", "web/js")
	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {
		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":     "Home Page",
				"framework": "Gin",
			},
		)

	})
	router.GET("/api/cart", middleware.HandleGuard, controller.GetCart)
	router.POST("/api/cart", middleware.HandleGuard, controller.AddCartProduct)
	router.PATCH("/api/cart", middleware.HandleGuard, controller.UpdateCartProduct)
	router.DELETE("/api/cart", middleware.HandleGuard, controller.DeleteCartProduct)
}
