package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(os.Getenv("HELLO_MESSAGE"))

	data, err := os.ReadFile("data/data.db")
	fmt.Println(string(data), err)

	time.Sleep(time.Second * 10)
	data, err = os.ReadFile("data/data.db")
	fmt.Println(string(data), err)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from docker server again")
	})

	r.Run(":8080")
}
