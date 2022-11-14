package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// gin.SetMode(gin.DebugMode)
	r.POST("/login", handleLogin)
	r.Run() // listen and serve on
}
