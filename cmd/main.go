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
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	servise := servise.New()

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
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
