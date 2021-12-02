package address

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Cidade   string `json:"cidade" validate:"required"`
	Uf       string `json:"uf" validate:"required"`
	ClientID uint
}

type AddressDTO struct {
	Cidade string `json:"cidade,omitempty"`
	Uf     string `json:"uf,omitempty"`
}
