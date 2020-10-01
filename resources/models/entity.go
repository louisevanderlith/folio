package models

import (
	"github.com/louisevanderlith/husk/validation"
	"github.com/louisevanderlith/kong/prime"
)

type Entity struct {
	Name           string
	ProfileKey     string
	User           prime.User
	Identification string
	Addresses      []Address
}

func (o Entity) Valid() error {
	return validation.Struct(o)
}
