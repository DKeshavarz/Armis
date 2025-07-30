package main

import (
	"log"
	"net/http"

	"github.com/DKeshavarz/armis/internal/commands"
	"github.com/DKeshavarz/armis/internal/server"
	"github.com/DKeshavarz/armis/internal/servise"
)

func main() {
	servise := servise.New()

	cmd := commands.New(servise)
	go cmd.Run()

	server := server.New(servise)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
