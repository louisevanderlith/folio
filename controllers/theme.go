package controllers

import (
	"errors"
	"net/http"

	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/folio/core"
)

type ThemeController struct {
	xontrols.APICtrl
}

//:site
func (c *ThemeController) Get() {
	profile := c.FindParam("site")

	if len(profile) == 0 {
		c.Serve(http.StatusNotFound, errors.New("invalid site"), nil)
		return
	}

	prof, err := core.GetProfileByName(profile)

	if err != nil {
		c.Serve(http.StatusNotFound, err, nil)
		return
	}

	theme := bodies.NewThemeSetting(prof.Title, c.Ctx().Host(), prof.ImageKey, "", prof.GTag)

	c.Serve(http.StatusOK, nil, theme)
}
