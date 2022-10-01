package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/probonopd/go-scribus"
	"github.com/sjustesen/scriflow/core/config"
	"github.com/sjustesen/scriflow/core/project"
)

func MountProjectRoutes(r *gin.RouterGroup) {

	// This route lists every scribus project
	// TODO: Restrict Access
	r.GET("list", func(c *gin.Context) {
		files := project.ListScribusProjects()
		if files != nil {
			c.JSON(http.StatusOK, files)
		} else {
			c.JSON(http.StatusOK, nil)
		}

	})

	// This route loads and parses a single Scribus file
	r.GET("load/:id", func(c *gin.Context) {
		doc, err := scribus.NewScribusDocumentFromFile(config.GetProjectPath() + "/doc1.sla")
		if err != nil {
			fmt.Println(err)
			fmt.Println("Couldn't load XML file")
		}
		c.XML(http.StatusOK, doc)
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.tmpl", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")

		if err != nil {
			log.Fatal(err)
			c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}

		filename := header.Filename
		out, _ := os.Create("public/" + filename)
		io.Copy(out, file)
		out.Close()

	})
}
