package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Profiles husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Profiles: husk.NewTable(Profile{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Profiles.Save()
}

func seed() {
	err := ctx.Profiles.Seed("db/profiles.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Profiles.Save()
}
