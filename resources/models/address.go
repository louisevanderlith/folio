package models

import "github.com/louisevanderlith/husk/validation"

type Address struct {
	StreetNo    int
	Street      string
	UnitNo      string
	EstateName  string
	Suburb      string
	City        string
	Province    string
	PostalCode  string
	Coordinates string
	IsDelivery  bool
}

func (o Address) Valid() error {
	return validation.Struct(o)
}
