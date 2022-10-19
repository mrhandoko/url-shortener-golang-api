package main

import (
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/handler"
	"url-shortener-golang-api/repository/sql"
	"url-shortener-golang-api/service"

	"github.com/gorilla/mux"
)

func main() {
	// Init DB Connection
	dbConfig := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "Ratna123",
		DB:         "urlshortener",
	}

	db, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}

	repo := sql.NewRepo(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/shorten/{url}", handler.ShortenURL).Methods("POST")
	router.HandleFunc("/stats/{id}", handler.GetStats).Methods("GET")
	router.HandleFunc("/{id}", handler.RedirectURL).Methods("GET")

	// Start server
	log.Println("Starting http server on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
