package main

// import (
// 	"io"
// 	"net/http"

// 	"github.com/dbzyuzin/swagger/internal/config"
// 	"github.com/dbzyuzin/swagger/internal/services"
// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )

// type Rest struct {
// 	lg      *zap.SugaredLogger
// 	service services.Service
// }

// func NewServer(lg *zap.SugaredLogger, cfg config.ServerConfig, service services.Service) *http.Server {
// 	if cfg.DebugMode {
// 		gin.SetMode(gin.DebugMode)
// 	} else {
// 		gin.SetMode(gin.ReleaseMode)
// 	}

// 	gin.DefaultWriter = io.Discard
// 	r := gin.Default()

// 	rest := Rest{lg, service}

// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"*"}
// 	config.AllowHeaders = []string{"*"}

// 	r.Use(cors.New(config))
// 	r.Use(func(ctx *gin.Context) {
// 		lg.Info("http request", ctx.Request.URL.Path)
// 	})
// 	r.POST("/users", rest.createUser)
// 	r.GET("/users/:name/exists", rest.userExists)

// 	return &http.Server{
// 		Addr:    cfg.ServerHost,
// 		Handler: r,
// 	}
// }
