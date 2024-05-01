package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name        string `json:"name" validate:"nonzero"`
	Description string `json:"description" validate:"nonzero"`
}

func ValidateCharacterData(character *Character) error {
	if err := validator.Validate(character); err != nil {
		return err
	}
	return nil
}
