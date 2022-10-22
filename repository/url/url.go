package url

import (
	"url-shortener-golang-api/repository"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) repository.RepositoryInterface {
	return &repo{
		db: db,
	}
}
