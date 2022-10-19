package main

import (
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	// Init DB Connection
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "Ratna123",
		DB:         "urlshortener",
	}

	connString := database.GetConnectionString(config)
	if err := database.Connect(connString); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/url/create", handler.GenerateURL)
	router.HandleFunc("/{uniqueCode}", handler.RedirectURL)

	// Start server
	log.Println("Starting http server on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
