package controllers

import (
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/mango/control"
)

type PortfolioController struct {
	control.APIController
}

func NewPortfolioCtrl(ctrlMap *control.ControllerMap) *PortfolioController {
	result := &PortfolioController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title CreatePortfolioItem
// @Description Creates a Portfolio Item on a current site
// @Param	body		body 	core.Portfolio	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *PortfolioController) Post() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	err = core.AddPortfolioSection(with.Key, with.Body.(core.Portfolio))

	req.Serve(nil, err)
}
