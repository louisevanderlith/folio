package core

import (
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
)

type Content struct {
	Realm       string
	Client      string
	LogoKey     *keys.TimeKey
	Language    string //en,af, en-US
	Banner      Banner
	SectionA    Section
	SectionB    Section
	Info        Information
	Colour      Colour
	Email       string
	Contacts    []Contact
	Description string
	GTag        string
}

func (o Content) Valid() error {
	return validation.Struct(o)
}
