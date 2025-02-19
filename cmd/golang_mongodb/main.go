package main

import (
	"context"
	"fmt"
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

	fmt.Println("Congif setup:", cfg)
	routers := routes.Routes()

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: routers,
	}
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)
	fmt.Println("Server Address:", cfg.Addr)

	go func() {
		slog.Info("server starting...") 
		err := server.ListenAndServe()
		slog.Info("server started") 
	
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed to start:", err)
			slog.Error("Failed to start server", slog.String("error", err.Error()))
			os.Exit(1)  // Explicitly exit to avoid confusion
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
