package rest

import (
	"github.com/dbzyuzin/swagger/internal/config"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users/:name/exists", userExists)

	return r
}
