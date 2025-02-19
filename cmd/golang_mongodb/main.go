package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	fmt.Println("Hello Mongodb")

// 	cfg := config.CongifInit()

// 	fmt.Println("Congif setup:", cfg)
// 	routers := routes.Routes()

// 	slog.Info("server started")
// 	server := http.Server{
// 		Addr:    cfg.Addr,
// 		Handler: routers,
// 	}
// 	done := make(chan os.Signal, 1)

// 	signal.Notify(done, os.Interrupt)

// 	go func() {
// 		err := server.ListenAndServe()

// 		if err != nil {
// 			log.Fatal("Failed to start server")
// 		}
// 	}()

// 	<-done

// 	slog.Info("Shutting down the server")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := server.Shutdown(ctx); err != nil {
// 		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
// 	}

// 	slog.Info("server shutdown")

// }

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func main() {
	http.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
