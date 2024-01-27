package main

import (
	"github.com/dbzyuzin/swagger/internal/transport/rest"
)

func main() {
	rest.NewServer().Run(":8080")
}
