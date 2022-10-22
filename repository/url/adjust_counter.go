package url

import (
	"log"
	"url-shortener-golang-api/entity"

	"gorm.io/gorm"
)

func (r *repo) AdjustCounter(id string) error {
	model := entity.URLshortener{}
	err := r.db.Transaction(func(tx *gorm.DB) error {
		errTx := tx.Model(model).Where("id = ?", id).Scan(&model).Error
		if errTx != nil {
			log.Fatalln(errTx.Error())
			return errTx
		}

		errTx = tx.Model(model).Where("id = ?", id).Update("redirect_count", model.RedirectCount+1).Error
		if errTx != nil {
			log.Fatalln(errTx.Error())
			return errTx
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}
