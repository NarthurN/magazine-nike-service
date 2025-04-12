package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"microservice1/internal/configs"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func initLoggerInfo() *log.Logger {
	log.Println("LoggerInfo initialized")
	return log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Llongfile)
}

func initLoggerError() *log.Logger {
	log.Println("LoggerError initialized")
	return log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Llongfile)
}

func initDb() *sql.DB {
	BDconfigs, err := configs.LoadDB()
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		BDconfigs.DBuser,
		BDconfigs.DBpass,
		//BDconfigs.DBhost,
		//BDconfigs.DBport,
		BDconfigs.DBtablename,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(context.Background())
	if err != nil {
		panic(err)
	}

	log.Println("Data Base initialized")
	return db
}

func InitServer(api *API) *http.Server {
	HTTPconfigs, err := configs.LoadHTTP()
	if err != nil {
		panic(err)
	}
	log.Println("http.Server initialized")
	return &http.Server{
		Addr:              HTTPconfigs.SRVaddress + ":" + HTTPconfigs.SRVport,
		Handler:           api.routes(),
		ReadTimeout:       HTTPconfigs.SRVreadTimeout,
		ReadHeaderTimeout: HTTPconfigs.SRVreadHeaderTimeout,
		WriteTimeout:      HTTPconfigs.SRVwriteTimeout,
		IdleTimeout:       HTTPconfigs.SRVidleTimeout,
	}
}
