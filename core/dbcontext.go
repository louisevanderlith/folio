package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Profiles husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Profiles: husk.NewTable(new(Profile)),
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
