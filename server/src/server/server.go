package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
)

const (
	userkey = "user"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Bootup() {
	fmt.Println("Running webserver...")
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.LoadHTMLGlob("./templates/*")

	mountRoutes(r)

	r.Run() // listen and serve o
}
