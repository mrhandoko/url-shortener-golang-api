package service

import (
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/repository"
)

type ServiceInterface interface {
	GenerateShortenURL(param entity.URLshortener) (string, error)
	RedirectURL(id string) (entity.URLshortener, error)
	GetURLs() ([]entity.URLshortener, error)
	GetURLByID(id string) (entity.URLshortener, error)
	UpdateURLByID(id string, shortenUrl entity.URLshortener) (string, error)
	DeleteURLByID(id string) (string, error)
}

type service struct {
	repository repository.RepositoryInterface
}

func NewService(repository repository.RepositoryInterface) ServiceInterface {
	return &service{repository: repository}
}
