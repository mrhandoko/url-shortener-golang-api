package url

import (
	"errors"
	"log"
	"url-shortener-golang-api/entity"
)

func (r *repo) UpdateURLByID(id string, url entity.URLshortener) error {
	result := r.db.Where("id = ?", id).Updates(&entity.URLshortener{
		OriginalURL:   url.OriginalURL,
		RedirectCount: url.RedirectCount,
	})
	if result.Error != nil {
		log.Fatalln(result.Error.Error())
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("update url failed")
	}

	return nil
}
