package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	scribus "github.com/probonopd/go-scribus"
	"github.com/sjustesen/scriflow/core/config"
)

func MountRoutes(r *gin.Engine) {

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
			systempath, err := os.Getwd()
			var arr []string
			dir, err := os.ReadDir(systempath + "Projects/")
			for _, file := range dir {
				arr = append(arr, file.Name())
				fmt.Println(arr)
			}
			if err == nil {
				json_dir, json_err := json.Marshal(arr)
				if json_err == nil {
					c.JSON(http.StatusOK, string(json_dir))
				} else {
					c.JSON(http.StatusBadRequest, json_err)
				}
			}
		})

		restricted.GET("load/:id", func(c *gin.Context) {
			systempath, err := os.Getwd()
			doc, err := scribus.NewScribusDocumentFromFile(systempath + config.GetProjectPath() + "/doc1.sla")
			if err != nil {
				fmt.Println(err)
				fmt.Println("Couldn't load XML file")
			}
			c.XML(http.StatusOK, doc)
		})

		restricted.GET("/upload", func(c *gin.Context) {
			c.HTML(http.StatusOK, "upload.tmpl", gin.H{})
		})

		restricted.POST("/upload", uploadfile)
	}
}
