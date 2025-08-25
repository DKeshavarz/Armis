// Copyright (c) 2025 Armis Dev Team
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/DKeshavarz/armis/internal/commands"
	"github.com/DKeshavarz/armis/internal/config"
	"github.com/DKeshavarz/armis/internal/logger"
	"github.com/DKeshavarz/armis/internal/server"
	"github.com/DKeshavarz/armis/internal/servise"
	"github.com/DKeshavarz/armis/internal/storage"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	
	mainLogger := logger.New("main")

	storage := storage.New(
		config.GetEnvAsBool("STORAGE_AUTO_SAVE", false), 
		config.GetEnvAsInt("STORAGE_SAVE_INTERVAL", 100), 
		config.GetEnv("STORAGE_FILE_PATH", "storage.json"),
	)
	
	servise := servise.New(storage)

	cmd := commands.New(servise)
	go func(){
		mainLogger.Info("starting command line ...")
		if err := cmd.Run(); err != nil {
			mainLogger.Error("error in command service", logger.Field{Key:"err", Value:err})
		}
	}()

	server := server.New(servise)
	port := config.GetEnv("PORT", "5432")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}
	go func() {
		mainLogger.Info("starting command web server...", logger.Field{Key:"Port", Value:port})
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mainLogger.Error("error in web server", logger.Field{Key:"err", Value:err})
		}
	}()
	
	<-ctx.Done()
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		mainLogger.Error("Server forced to shutdown", logger.Field{Key:"err", Value:err})
	}

	if err := storage.Close(); err != nil {
		mainLogger.Error("Storage forced to shutdown", logger.Field{Key:"err", Value:err})
	}

	mainLogger.Info("shut down")
}
