package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello API")
	})
	http.ListenAndServe(":9234", nil)
}
