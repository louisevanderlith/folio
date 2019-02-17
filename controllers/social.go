package controllers

import (
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango/folio/core"
)

type SocialController struct {
	control.APIController
}

func NewSocialCtrl(ctrlMap *control.ControllerMap) *SocialController {
	result := &SocialController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title CreateSocialLink
// @Description Creates a Social Link on a current site
// @Param	body		body 	core.SocialLink	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *SocialController) Post() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	err = core.AddSocialLink(with.Key, with.Body.(core.SocialLink))

	req.Serve(nil, err)
}
