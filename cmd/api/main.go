package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dbzyuzin/swagger/internal/config"
	"github.com/dbzyuzin/swagger/internal/repostories/memory"
	"github.com/dbzyuzin/swagger/internal/services"
	"github.com/dbzyuzin/swagger/internal/transport/rest"
)

func main() {
	repo := &memory.Repository{}
	service := services.New(repo)

	server := rest.NewServer(config.ServerHost, service)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Panicln("Shutdown error:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panic(err)
	}
}
