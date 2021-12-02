package client

import (
	"go-atividade-a3/address"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Nome     string          `json:"nome" validate:"required"`
	Endereco address.Address `json:"endereco" validate:"required" gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ClientDTO struct {
	ID       uint                `json:"id"`
	Nome     string              `json:"nome"`
	Endereco *address.AddressDTO `json:"endereco,omitempty"`
}

type ClientsByCityDTO struct {
	Cidade   string      `json:"cidade"`
	Uf       string      `json:"uf"`
	Clientes []ClientDTO `json:"clientes"`
}
