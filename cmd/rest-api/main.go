package main

import (
	"go-to-school/main/entrypoints"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "api", log.Flags())
	helloHandler := entrypoints.NewHello(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9234", serveMux)
}
