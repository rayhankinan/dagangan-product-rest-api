package main

import (
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/middlewares"
	"dagangan-product-rest-api/routes"
	databaseService "dagangan-product-rest-api/services/database"
)

func main() {
	// Configure Runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Configure Database
	databaseService.DB.GetConnection()

	// Gin Framework
	engine := gin.Default()

	// Middlewares
	engine.Use(middlewares.CORSMiddleware())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Routes
	routes.UserRoute(engine)
	routes.ProductRoute(engine)
	routes.NotFoundRoute(engine)

	// Run App
	engine.Run()
}
