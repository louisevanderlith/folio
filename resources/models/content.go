package models

import (
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/validation"
)

type Content struct {
	Profile  string
	Language string //en,af, en-US
	Banner   Banner
	SectionA Section
	SectionB Section
	Info     Information
	Colour   Colour
}

func (c Content) Valid() error {
	return validation.Struct(c)
}

type Section struct {
	Heading  string
	Text     string
	ImageUrl string
	ImageKey keys.TimeKey
}

type Information struct {
	Heading string
	Text    string
	Blocks  []SimpleBlock
}

type SimpleBlock struct {
	Icon string
	Text string
}

type Banner struct {
	Background keys.TimeKey
	Image      keys.TimeKey
	Heading    string
	Subtitle   string
}

type Colour struct {
	Primary    RGB
	Secondary  RGB
	Tertiary   RGB
	Shadows    RGB
	Accent     RGB
	Background RGB
}

type RGB struct {
	Red   int
	Green int
	Blue  int
}
