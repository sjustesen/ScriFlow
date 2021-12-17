package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/sjustesen/scriflow/core/config"

	"github.com/gin-gonic/gin"
	"github.com/probonopd/go-scribus"

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

	// -- Setup Routes --
	// TODO: getting quite lenghty - perhaps split up at some point?

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	api := r.Group("/api")
	api.Use(AuthRequired)
	{
		api.GET("loadfile", profile)
	}

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

		restricted.GET("list", func(c *gin.Context) {
			c.JSON(http.StatusOK, "list projects")
		})

		restricted.GET("load/:id", func(c *gin.Context) {
			doc, err := scribus.NewScribusDocumentFromFile(config.GetProjectPath() + "Document-1.sla")
			if err != nil {
				fmt.Println("Couldn't load XML file")
			}
			c.XML(http.StatusOK, doc)
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
	// todo
}
