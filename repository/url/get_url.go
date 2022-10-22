package url

import (
	"errors"
	"log"
	"url-shortener-golang-api/entity"
)

func (r *repo) GetURLs() (urls []entity.URLshortener, err error) {
	if err := r.db.Find(&urls).Error; err != nil {
		log.Fatalln(err.Error())
		return urls, err
	}

	return urls, nil
}

func (r *repo) GetURLByID(id string) (entity.URLshortener, error) {
	var url entity.URLshortener
	result := r.db.Where("id = ? ", id).Find(&url)
	if result.Error != nil {
		err := result.Error
		log.Fatalln(err.Error())
		return url, err
	}

	if result.RowsAffected < 1 {
		return url, errors.New("url not found")
	}

	return url, nil
}
