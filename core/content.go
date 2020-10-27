package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
)

type Content struct {
	Realm    string
	Client   string
	LogoKey  hsk.Key
	Language string //en,af, en-US
	Banner   Banner
	SectionA Section
	SectionB Section
	Info     Information
	Colour   Colour
	Contacts []Contact
}

func NewContent() Content {
	return Content{
		Realm:    "",
		Client:   "",
		LogoKey:  keys.CrazyKey(),
		Language: "",
		Banner: Banner{
			Background: keys.CrazyKey(),
			Image:      keys.CrazyKey(),
			Heading:    "",
			Subtitle:   "",
		},
		SectionA: Section{},
		SectionB: Section{},
		Info:     Information{},
		Colour:   Colour{},
		Contacts: nil,
	}
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
