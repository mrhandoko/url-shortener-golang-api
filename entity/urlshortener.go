package entity

type URLshortener struct {
	ID            int32  `json:"id"`
	UniqueCode    string `json:"unique_code"`
	URL           string `json:"url"`
	OriginalURL   string `json:"original_url"`
	RedirectCount int32  `json:"redirect_count"`
}

func (URLshortener) TableName() string {
	return "generated_url"
}
