package url

import (
	"log"
	"url-shortener-golang-api/entity"
)

func (r *repo) CreateURL(params *entity.URLshortener) error {
	if err := r.db.Create(&params).Error; err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}
