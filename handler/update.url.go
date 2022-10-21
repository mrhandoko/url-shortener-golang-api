package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/entity"

	"github.com/gorilla/mux"
)

func UpdateURLByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID := params["id"]

	body, _ := ioutil.ReadAll(r.Body)
	var url entity.URLshortener
	json.Unmarshal(body, &url)

	result := database.Connector.Debug().Table(url.TableName()).
		Select("original_url", "redirect_count").
		Where("id = ?", ID).
		Update(&url)

	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}
	// json.NewEncoder(w).Encode(url)
}

func UpdateURLByUniqueCode(w http.ResponseWriter, r *http.Request) {
	panic("implement me!")
}
