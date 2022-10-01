package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjustesen/scriflow/core/controllers"
)

func mountRoutes(r *gin.Engine) {

	// Public user home route
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	restrictedRouteGroup := r.Group("/projects")
	restrictedRouteGroup.Use(AuthRequired)
	{
		controllers.MountProjectRoutes(restrictedRouteGroup)
	}

	backendusers := r.Group("/users")
	backendusers.Use(AuthRequired)
	{
		controllers.MountUsersRoutes(restrictedRouteGroup)

	}

}

func AuthRequired(c *gin.Context) {

}
