package rest

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users/:name/exists", userExists)

	return r
}
