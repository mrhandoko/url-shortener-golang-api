package sql

import (
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/repository"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) repository.RepoInterface {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(params *entity.URLshortener) error {
	err := r.db.Create(&params).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) GetByID(id string) (*entity.URLshortener, error) {
	model := entity.URLshortener{}
	err := r.db.Model(model).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *repo) AdjustCounter(id string) error {
	model := entity.URLshortener{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		errTx := tx.Model(model).Where("id = ?", id).Scan(&model).Error
		if errTx != nil {
			return errTx
		}

		errTx = tx.Model(model).Where("id = ?", id).Update("redirect_counter", model.RedirectCount+1).Error
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil

}
