package main

import (
	"net/http"

	"github.com/dbzyuzin/swagger/models"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.User{
			Name: "Dima",
		})
	})

	r.Run(":8080")
}
