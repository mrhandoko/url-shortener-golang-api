package main

import (
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/handler"
	"url-shortener-golang-api/repository/url"
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

	repository := url.NewRepo(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/urls/create", handler.GenerateURL).Methods("POST")
	router.HandleFunc("/{id}", handler.RedirectURL).Methods("GET")
	router.HandleFunc("/api/urls", handler.GetURLData).Methods("GET")
	router.HandleFunc("/api/url/{id}", handler.GetURLByID).Methods("GET")
	router.HandleFunc("/api/url/{id}", handler.UpdateURLByID).Methods("PUT")
	router.HandleFunc("/api/url/{id}", handler.DeleteURLByID).Methods("DELETE")

	// Start server
	log.Println("Starting http server on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}
