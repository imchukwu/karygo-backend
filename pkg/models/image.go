package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Title   string `json:"title"`
	AltText uint   `json:"alt_text"`
	Url     string `json:"url"`
}
