package server


import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/sessions"
)

const (
  userkey = "user"
)

func Bootup() {
    fmt.Println("Running webserver...")

    r := gin.Default()
    r.LoadHTMLGlob("./templates/*")

    r.GET("/", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{

      })
    })

    restricted := r.Group("/user")
    restricted.Use(AuthRequired) 
    {
      restricted.GET("/upload", uploadfile)
      restricted.GET("/me", profile)
    }
    
    r.Run() // listen and serve o
}

func uploadfile(c *gin.Context) {
}

func profile(c *gin.Context) {
}

func AuthRequired(c *gin.Context) {

}
