package rest

import (
	"github.com/dbzyuzin/swagger/internal/config"
	"github.com/dbzyuzin/swagger/internal/services"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	service services.Service
}

func NewServer(service services.Service) *gin.Engine {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	rest := Rest{service}

	r.POST("/users", rest.createUser)
	r.GET("/users/:name/exists", rest.userExists)

	return r
}
