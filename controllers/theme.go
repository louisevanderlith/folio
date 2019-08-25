package controllers

import (
	"errors"
	"net/http"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/folio/core"
)

type Theme struct {
}

//:site
func (c *Theme) Get(ctx context.Contexer) (int, interface{}) {
	profile := ctx.FindParam("site")

	if len(profile) == 0 {
		return http.StatusBadRequest, errors.New("invalid site")
	}

	prof, err := core.GetProfileByName(profile)

	if err != nil {
		return http.StatusNotFound, err
	}

	theme := bodies.NewThemeSetting(prof.Title, ctx.Host(), prof.ImageKey, "", prof.GTag)

	return http.StatusOK, theme
}
