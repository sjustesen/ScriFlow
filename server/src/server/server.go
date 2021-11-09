package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/sessions"
)

const (
	userkey = "user"
)

func Bootup() {
	fmt.Println("Running webserver...")
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	// -- Setup Routes --
	// TODO: getting quite lenghty - perhaps split up at some point?

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	users := r.Group("/users")
	users.Use(AuthRequired)
	{
		users.GET("/me", profile)
	}

	restricted := r.Group("/projects")
	restricted.Use(AuthRequired)
	{
		restricted.GET(":id", func(c *gin.Context) {
			c.HTML(http.StatusOK, "loadproject.tpml", gin.H{})
		})

		restricted.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload.tmpl", gin.H{})
		})

		restricted.POST("/upload", uploadfile)
	}

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

}
