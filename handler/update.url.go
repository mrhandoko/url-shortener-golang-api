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

	result := database.Connector.Table(url.TableName()).
		Where("id = ?", ID).
		Update(map[string]interface{}{
			"original_url":   url.OriginalURL,
			"redirect_count": url.RedirectCount,
		})

	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}

	result.Scan(&url)
	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode(url)
	} else {
		json.NewEncoder(w).Encode("Update url failed")
	}

}

func UpdateURLByUniqueCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	uniqueCode := params["uniqueCode"]

	body, _ := ioutil.ReadAll(r.Body)
	var url entity.URLshortener
	json.Unmarshal(body, &url)

	result := database.Connector.Debug().Table(url.TableName()).
		Where("unique_code = ?", uniqueCode).
		Update(map[string]interface{}{
			"original_url":   url.OriginalURL,
			"redirect_count": url.RedirectCount,
		})

	if result.Error != nil {
		log.Fatalln(result.Error.Error())
	}

	result.Scan(&url)
	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode(url)
	} else {
		json.NewEncoder(w).Encode("Update url failed")
	}
}
