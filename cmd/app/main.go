package main

import (
	"context"
	"microservice1/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	api := server.NewAPI()
	defer api.Magazines.DB.Close()
	defer api.Broker.Close()
	defer api.Cache.Close()
	
	server := server.InitServer(api)

	// Graceful Shutdown
	serverErrors := make(chan error, 1)

	go func() {
		api.LogInfo.Println("Server is listening on port :8080 ...")
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		api.LogError.Fatalf("Error starting server: %v", err)
	case <-shutdown:
		api.LogInfo.Println("Starting graceful shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			api.LogError.Printf("Graceful shutdown failed: %v", err)
			if err := server.Close(); err != nil {
				api.LogError.Fatalf("Forced shutdown failed: %v", err)
			}
		}
		api.LogInfo.Println("Server stopped")
	}
}
