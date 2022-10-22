package service

import (
	"errors"
	"fmt"
	"strings"
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/helper"
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
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return &service{repo: repo}
}

func (s *service) GenerateShortenURL(param entity.URLshortener) (string, error) {
	if helper.ValidParam(param) {
		fmt.Println(param)
		return "", errors.New("invalid param")
	}

	id := helper.GenerateUniqueCode(6)
	if err := s.repo.CreateURL(&entity.URLshortener{
		ID:          id,
		ShortURL:    fmt.Sprintf(`%s/%s`, entity.URL, id),
		OriginalURL: param.OriginalURL,
	}); err != nil {
		return "", err
	}

	return id, nil
}

func (s *service) RedirectURL(id string) (data entity.URLshortener, err error) {
	if strings.EqualFold(id, "") {
		return data, errors.New("invalid short url")
	}

	data, err = s.repo.GetURLByID(id)
	if err != nil {
		return data, err
	}

	err = s.repo.AdjustCounter(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) GetURLs() (urls []entity.URLshortener, err error) {
	urls, err = s.repo.GetURLs()
	if err != nil {
		return urls, err
	}

	return urls, nil
}

func (s *service) GetURLByID(id string) (data entity.URLshortener, err error) {
	if strings.EqualFold(id, "") {
		return data, errors.New("invalid id param")
	}

	data, err = s.repo.GetURLByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) UpdateURLByID(id string, url entity.URLshortener) (string, error) {
	if strings.EqualFold(id, "") {
		return "", errors.New("invalid id param")
	}
	if strings.EqualFold(url.OriginalURL, "") {
		fmt.Println(url)
		return "", errors.New("invalid original url")
	}

	err := s.repo.UpdateURLByID(id, url)
	if err != nil {
		return "", err
	}

	return "Update Success", nil
}

func (s *service) DeleteURLByID(id string) (string, error) {
	if strings.EqualFold(id, "") {
		return "", errors.New("invalid id param")
	}

	err := s.repo.DeleteURLByID(id)
	if err != nil {
		return "", err
	}

	return "Delete Success", nil
}
