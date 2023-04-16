package main

import (
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/middlewares"
	databaseService "dagangan-product-rest-api/services/database"
)

func main() {
	// Configure Runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Initialize Database
	databaseService.DB.GetConnection()

	// Gin Framework
	engine := gin.Default()

	// Middlewares
	engine.Use(middlewares.CORSMiddleware())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Run App
	engine.Run()
}
