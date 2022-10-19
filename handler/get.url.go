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
