package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"main.go/internal/config"
	"main.go/routes"
)


func main() {
	fmt.Println("Hello Mongodb")

	cfg := config.CongifInit()
	// rdctx := context.Background()

	// if err := c.Ping(rdctx); err != nil {
	// 	log.Panic("Failed to connect Redis")
	// }
	// log.Println("Redis Connected")
	// if err := c.Set(rdctx, "user:name", "Jatin", 0); err != nil {
	// 	log.Println("Error: Value is not stored")
	// }

	// log.Println("Value stored in redis")

	// res, err := c.Get(rdctx, "user:name")
	// if err != nil {
	// 	log.Println("Error: Value is not stored", err.Error())
	// }
	// fmt.Println("Redis value:", res)

	fmt.Println("Congif setup:", cfg)
	routers := routes.Routes()

	slog.Info("server started")
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: routers,
	}
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown")

}
