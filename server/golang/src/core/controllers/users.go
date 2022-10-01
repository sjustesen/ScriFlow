package controllers

import "github.com/gin-gonic/gin"

func MountUsersRoutes(r *gin.RouterGroup) {

	r.GET("/me", func(c *gin.Context) {
		// get user specific fields if he/she is logged in
	})
}
