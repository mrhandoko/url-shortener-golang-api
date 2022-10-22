package entity

import "time"

type URLshortener struct {
	ID            string `json:"id"`
	ShortURL      string `json:"short_url"`
	OriginalURL   string `json:"original_url"`
	RedirectCount int32  `json:"redirect_count"`
	CreatedAt     time.Time
}

func (URLshortener) TableName() string {
	return "generated_url"
}

const (
	URL = "http://localhost:8090"
)
