package main

import (
	"fmt"
	"log"
	"net/http"
)

const name = "UoF"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Kubernetes ♡ %s!", name)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
