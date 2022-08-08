package models

import (
	"gorm.io/gorm"
)

type Lga struct {
	gorm.Model
	Name    string `json:"name"`
	StateID uint   `json:"stateId"`
	State   *State `json:"state" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
