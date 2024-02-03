package rest

import (
	"net/http"

	"github.com/dbzyuzin/swagger/internal/config"
	"github.com/dbzyuzin/swagger/internal/services"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	service services.Service
}

func NewServer(host string, service services.Service) *http.Server {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	rest := Rest{service}

	r.POST("/users", rest.createUser)
	r.GET("/users/:name/exists", rest.userExists)

	return &http.Server{
		Addr:    host,
		Handler: r,
	}
}
