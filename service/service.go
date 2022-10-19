package service

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/helper"
	"url-shortener-golang-api/repository"
)

type service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	GenerateShortenURL(param entity.URLshortener) (string, error)
	GetDataByID(id string) (*entity.URLshortener, error)
	Redirect(id string) (*entity.URLshortener, error)
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &service{repo: repo}
}

func (s *service) GenerateShortenURL(param entity.URLshortener) (string, error) {
	if !validParam(param) {
		return "", errors.New("invalid parameter")
	}
	id := helper.GenerateUniqueCode(6)
	err := s.repo.Create(&entity.URLshortener{
		ID:          id,
		OriginalURL: param.OriginalURL,
		ShortURL:    fmt.Sprintf(`%s/%s`, entity.URL, id),
		CreatedAt:   time.Now(),
	})

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *service) GetDataByID(id string) (*entity.URLshortener, error) {
	if strings.EqualFold(id, "") {
		return nil, errors.New("invalid short url")
	}

	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) Redirect(id string) (*entity.URLshortener, error) {
	if strings.EqualFold(id, "") {
		return nil, errors.New("invalid short url")
	}

	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	err = s.repo.AdjustCounter(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func validParam(param entity.URLshortener) bool {
	return !strings.EqualFold(param.OriginalURL, "")
}
