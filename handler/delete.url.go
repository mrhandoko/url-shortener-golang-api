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

func DeleteURLByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID := params["id"]

	body, _ := ioutil.ReadAll(r.Body)
	var url entity.URLshortener
	json.Unmarshal(body, &url)

	result := database.Connecto.Table(url.TableName()).
		Where("id = ?", ID).
		Delete(&url)

	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}

	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode("Delete success")
	} else {
		json.NewEncoder(w).Encode("Delete url failed")
	}
}

func DeleteURLByUniqueCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	uniqueCode := params["uniqueCode"]

	body, _ := ioutil.ReadAll(r.Body)
	var url entity.URLshortener
	json.Unmarshal(body, &url)

	result := database.Connector.Table(url.TableName()).
		Where("uniqu_code = ?", uniqueCode).
		Delete(&url)

	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}

	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode("Delete success")
	} else {
		json.NewEncoder(w).Encode("Delete url failed")
	}
}
