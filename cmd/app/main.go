package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type API struct {
	logInfo  *log.Logger
	logError *log.Logger
	db       *sql.DB
	server   *http.Server
}

func newAPI(logInfo, logError *log.Logger, db *sql.DB, server *http.Server) *API {
	return &API{
		logInfo:  logInfo,
		logError: logError,
		db:       db,
		server:   server,
	}
}

func main() {
	r := http.NewServeMux()
	//r.Handle("/getMagazinesByCity", nil)

	loggerInfo := log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Llongfile)
	loggerError := log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Llongfile)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ErrorLog:          loggerError,
		ReadTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 1 * time.Minute,
		WriteTimeout:      1 * time.Minute,
		IdleTimeout:       5 * time.Minute,
	}

	api := newAPI(loggerInfo, loggerError, nil, server)

	// Graceful Shutdown
	serverErrors := make(chan error, 1)

	go func() {
		api.logInfo.Println("Server is listening on port :8080 ...")
		serverErrors <- api.server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		api.logError.Fatalf("Error starting server: %v", err)
	case <-shutdown:
		api.logInfo.Println("Starting graceful shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			api.logError.Printf("Graceful shutdown failed: %v", err)
			if err := server.Close(); err != nil {
				api.logError.Fatalf("Forced shutdown failed: %v", err)
			}
		}
		api.logInfo.Println("Server stopped")
	}
}
