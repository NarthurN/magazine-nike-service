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
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

func initLoggerInfo() *log.Logger {
	log.Println("LoggerInfo initialized")
	return log.New(os.Stdout, "INFO", log.Ldate|log.Ltime|log.Llongfile)
}

func initLoggerError() *log.Logger {
	log.Println("LoggerError initialized")
	return log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Llongfile)
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	if pong == "PONG" {
		log.Println("Redis initialized")
	}

	return rdb
}

func initRabitMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	log.Println("RabbitMQ initialized")
	return ch
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
