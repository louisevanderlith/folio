package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type Banner struct {
	Background hsk.Key
	Image      hsk.Key `hsk:"null"`
	Heading    string
	Subtitle   string
}
