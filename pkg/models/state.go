package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name string `json:"name"`
}
