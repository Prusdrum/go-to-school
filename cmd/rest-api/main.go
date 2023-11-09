package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Request received")
		body, err := io.ReadAll(req.Body)

		if err != nil {
			http.Error(res, "Error ocurred", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(res, "Hello %s", body)
	})
	http.ListenAndServe(":9234", nil)
}
