package main

import (
	"github.com/dbzyuzin/swagger/internal/repostories/memory"
	"github.com/dbzyuzin/swagger/internal/services"
	"github.com/dbzyuzin/swagger/internal/transport/rest"
)

func main() {
	repo := &memory.Repository{}
	service := services.New(repo)

	rest.NewServer(service).Run(":8080")
}
