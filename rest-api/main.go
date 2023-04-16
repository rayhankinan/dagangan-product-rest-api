package main

import (
	"runtime"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"dagangan-product-rest-api/middlewares"
)

func main() {
	// Configure runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Gin Framework
	engine := gin.Default()

	// Middlewares
	engine.Use(middlewares.CORSMiddleware())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Run App
	engine.Run()
}
