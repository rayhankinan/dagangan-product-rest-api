package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"runtime"
)

func main() {
	// Configure runtime
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Gin Framework
	engine := gin.Default()

	// Middlewares
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	// Run App
	engine.Run()
}
