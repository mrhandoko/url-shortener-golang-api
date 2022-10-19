package repository

import (
	"url-shortener-golang-api/entity"
)

type RepoInterface interface {
	Create(*entity.URLshortener) error
	GetByID(id string) (*entity.URLshortener, error)
	AdjustCounter(shortURL string) error
}
