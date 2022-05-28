package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"server/rest"
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

	rest.MountRoutes(r)

	r.Run() // listen and serve o
}

func uploadfile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
		log.Fatal(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

}

func profile(c *gin.Context) {
}

func AuthRequired(c *gin.Context) {
	// todo
}
