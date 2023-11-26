package main

import (
	api "awesomeProject"
	"log"
	"net/http"
)

func main() {
	api := api.New("localhost:8080", http.NewServeMux())
	api.FillEndPoints()
	log.Fatal(api.ListenAndServe())
}
