package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/DKeshavarz/armis/internal/commands"
	"github.com/DKeshavarz/armis/internal/config"
	"github.com/DKeshavarz/armis/internal/server"
	"github.com/DKeshavarz/armis/internal/servise"
	"github.com/DKeshavarz/armis/internal/storage"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	storage := storage.New(true, 100, "storage.json")
	servise := servise.New(storage)

	cmd := commands.New(servise)
	go func(){
		if err := cmd.Run(); err != nil {
			log.Printf("error in command service: %s", err)
		}
	}()

	server := server.New(servise)
	srv := &http.Server{
		Addr:    ":" + config.GetFromEnv("PORT"),
		Handler: server,
	}
	go func() {
		log.Printf("web server started on %s", config.GetFromEnv("PORT"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	
	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %s", err.Error())
	}

	if err := storage.Close(); err != nil {
		log.Fatalf("Storage shutdown with error %s", err.Error())
	}

	log.Println("grasfully shutdwn")
}
