package url

import (
	"errors"
	"log"
	"url-shortener-golang-api/entity"
)

func (r *repo) DeleteURLByID(id string) error {
	var url *entity.URLshortener
	result := r.db.Where("id = ?", id).Delete(&url)
	if result.Error != nil {
		log.Fatalln(result.Error.Error())
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("delete url failed")
	}

	return nil
}
