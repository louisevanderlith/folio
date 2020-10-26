package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
)

type Content struct {
	Realm    string
	Client   string
	LogoKey  keys.TimeKey
	Language string //en,af, en-US
	Banner   Banner
	SectionA Section
	SectionB Section
	Info     Information
	Colour   Colour
	Contacts []Contact
}

func (o Content) Valid() error {
	return validation.Struct(o)
}

func (o Content) Create() (hsk.Key, error) {
	return ctx.Content.Create(o)
}

func (o Content) Update(key hsk.Key) error {
	return ctx.Content.Update(key, o)
}
