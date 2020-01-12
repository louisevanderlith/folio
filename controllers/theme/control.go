package theme

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/folio/core"
	"net/http"
)

func Get(c *gin.Context) {
	profile := c.Param("name")

	if len(profile) == 0 {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid site"))
	}

	prof, err := core.GetProfileByName(profile)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	theme := element.NewIdentity(prof.Title, c.Request.Host, prof.ImageKey, "", prof.GTag)

	c.JSON(http.StatusOK, theme)
}
