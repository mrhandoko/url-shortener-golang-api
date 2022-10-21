package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"url-shortener-golang-api/database"
	"url-shortener-golang-api/entity"

	"github.com/gorilla/mux"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uniqueCode := params["uniqueCode"]

	var generatedUrl entity.URLshortener
	db := database.Connector.Table(generatedUrl.TableName())
	tx := db.Begin()

	if err := tx.Error; err != nil {
		log.Fatalln(err.Error())
		tx.Rollback()
	}

	if err := tx.Where("unique_code = ?", uniqueCode).Scan(&generatedUrl).Error; err != nil {
		log.Fatalln(err.Error())
		tx.Rollback()
	}

	result := tx.Where("unique_code = ?", uniqueCode).Update("redirect_count", generatedUrl.RedirectCount+1)
	if result.Error != nil {
		log.Fatalln(result.Error.Error())
		tx.Rollback()
	}

	tx.Commit()
	if result.RowsAffected > 0 {
		http.Redirect(w, r, generatedUrl.OriginalURL, http.StatusSeeOther)
		json.NewEncoder(w).Encode("Redirect success")
	}
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var urls []entity.URLshortener

	if err := database.Connector.Find(&urls).Error; err != nil {
		log.Fatalln(err.Error())
	}
	json.NewEncoder(w).Encode(urls)
}

func GetURLByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	ID := params["id"]
	var url entity.URLshortener

	if err := database.Connector.Find(&url, ID).Error; err != nil {
		log.Fatalln(err.Error())
	}
	json.NewEncoder(w).Encode(url)
}

func GetURLByUniqueCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	UniqueCode := params["uniqueCode"]
	var url entity.URLshortener

	if err := database.Connector.First(&url, "unique_code = ?", UniqueCode).Error; err != nil {
		log.Fatalln(err.Error())
	}

	json.NewEncoder(w).Encode(url)
}
