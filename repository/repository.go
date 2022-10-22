package repository

import (
	"url-shortener-golang-api/entity"
)

type RepositoryInterface interface {
	AdjustCounter(id string) error

	CreateURL(*entity.URLshortener) error
	GetURLs() ([]entity.URLshortener, error)
	GetURLByID(id string) (entity.URLshortener, error)
	UpdateURLByID(id string, url entity.URLshortener) error
	DeleteURLByID(id string) error
}
