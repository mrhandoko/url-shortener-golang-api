package helper

import (
	"strings"
	"url-shortener-golang-api/entity"
)

func ValidParam(param entity.URLshortener) bool {
	return strings.EqualFold(param.OriginalURL, "")
}
