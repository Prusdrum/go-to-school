package main

import (
	"context"
	"go-to-school/main/entrypoints"
	"go-to-school/main/infrastructure/tracing"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "api", log.Flags())
	helloHandler := entrypoints.NewHello(logger)
	groupsHandler := entrypoints.NewGroups(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/groups", groupsHandler)

	traced := tracing.NewTracingHandler(logger, serveMux)

	server := &http.Server{
		Addr:         ":9234",
		Handler:      traced,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// wrap with go func to not block
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// to gracefully shutdown.
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	logger.Println("Received terminate, graceful shutdown", signal)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
