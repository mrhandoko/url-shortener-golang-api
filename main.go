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
	router.HandleFunc("/api/urls", handler.GetURL)
	router.HandleFunc("/api/url/{id}", handler.GetURLByID).Methods("GET")
	router.HandleFunc("/api/url/unique/{uniqueCode}", handler.GetURLByUniqueCode).Methods("GET")
	router.HandleFunc("/api/url/{id}", handler.UpdateURLByID).Methods("PUT")
	router.HandleFunc("/api/url/unique/{uniqueCode}", handler.UpdateURLByUniqueCode).Methods("PUT")
	router.HandleFunc("/api/url/{id}", handler.DeleteURLByID).Methods("DELETE")
	router.HandleFunc("/api/url/unique/{uniqueCode}", handler.DeleteURLByID).Methods("DELETE")

	// Start server
	log.Println("Starting http server on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
