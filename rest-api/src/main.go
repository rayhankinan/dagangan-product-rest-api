package src

import (
	"github.com/gin-gonic/gin"

	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Gin Framework
	engine := gin.Default()

	// Run App
	engine.Run()
}
