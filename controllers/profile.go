package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/mango/control"
)

type ProfileController struct {
	control.APIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap) *ProfileController {
	result := &ProfileController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /all/:pagesize [get]
func (req *ProfileController) Get() {
	page, size := req.GetPageData()
	results := core.GetProfiles(page, size)

	req.Serve(results, nil)
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {core.Profile} core.Profile
// @router /:site [get]
func (req *ProfileController) GetOne() {
	siteParam := req.Ctx.Input.Param(":site")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		byName, err := core.GetProfileByName(siteParam)
		req.Serve(byName, err)
		return
	}

	result, err := core.GetProfile(key)

	req.Serve(result, err)
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ProfileController) Post() {
	var site core.Profile
	json.Unmarshal(req.Ctx.Input.RequestBody, &site)

	rec := site.Create()

	req.Serve(rec, nil)
}

// @Title UpdateWebsite
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ProfileController) Put() {
	with, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	body := with.Body.(core.Profile)
	err = body.Update(with.Key)

	req.Serve(nil, err)
}
