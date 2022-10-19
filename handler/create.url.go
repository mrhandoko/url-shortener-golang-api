package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/helper"
)

func GenerateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var generatedUrl entity.URLshortener
	json.Unmarshal(body, &generatedUrl)

	randUniqueCode := helper.GenerateUniqueCode(6)
	generatedUrl.UniqueCode = randUniqueCode
	generatedUrl.URL = "http://localhost:8090/" + randUniqueCode

	if err := database.Connector.Table(generatedUrl.TableName()).Create(&generatedUrl).Error; err != nil {
		log.Fatalln(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(generatedUrl)
	}
}
