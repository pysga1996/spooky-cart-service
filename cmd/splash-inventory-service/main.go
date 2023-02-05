package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thanh-vt/splash-inventory-service/internal"
	"github.com/thanh-vt/splash-inventory-service/internal/config"
	"github.com/thanh-vt/splash-inventory-service/internal/controller"
	customMiddleware "github.com/thanh-vt/splash-inventory-service/internal/middleware"
	"html/template"
	"net/http"
	"os"
)

type HtmlModel = map[string]interface{}

func main() {
	internal.HttpClient = new(http.Client)
	err := config.ConnectRedis() // connect redis
	if err != nil {
		panic(err.Error())
	}
	err = config.ConnectDatabase() // connect database
	if err != nil {
		panic(err.Error())
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger, middleware.Recoverer)
	router.Use(customMiddleware.HandleError, customMiddleware.HandleToken)

	registerRoutes(router)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	// Start serving the application
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		panic(err.Error())
	}

}

func registerRoutes(router *chi.Mux) {
	// Set the router as the default one provided by Gin

	router.Handle("/css/*",
		http.StripPrefix("/css", http.FileServer(http.Dir("web/css"))))
	router.Handle("/js/*",
		http.StripPrefix("/js", http.FileServer(http.Dir("web/js"))))
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	var templates = template.Must(template.ParseGlob("web/templates/*"))
	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			model := HtmlModel{
				"title":     "Home Page",
				"framework": "Chi",
			}
			if err := templates.Lookup("index.html").Execute(w, model); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	})
	router.Route("/revice-commerce", func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Route("/supplier", func(r chi.Router) {
				r.Use(customMiddleware.HandleGuard)
				r.Get("/", controller.GetAllSupplier)  // GET /supplier
				r.Post("/", controller.CreateSupplier) // POST /supplier/123
				r.Route("/{code}", func(r chi.Router) {
					r.Get("/", controller.GetSupplier)          // GET /supplier/123
					r.Put("/", controller.UpdateCartProduct)    // PUT /supplier/123
					r.Delete("/", controller.DeleteCartProduct) // DELETE /supplier/123
				})
			})
		})
		//
		//// GET /articles/whats-up
		//r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	})

}
