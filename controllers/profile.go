package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/folio/core"
)

type ProfileController struct {
	xontrols.APICtrl
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /all/:pagesize [get]
func (req *ProfileController) Get() {
	page, size := req.GetPageData()
	results := core.GetProfiles(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {core.Profile} core.Profile
// @router /:site [get]
func (req *ProfileController) GetOne() {
	siteParam := req.Ctx.FindParam("site")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		byName, err := core.GetProfileByName(siteParam)

		if err != nil {
			req.Serve(http.StatusNotFound, err, nil)
			return
		}

		req.Serve(http.StatusOK, nil, byName)
		return
	}

	result, err := core.GetProfile(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *ProfileController) Post() {
	var site core.Profile
	err := req.Ctx.Body(&site)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec := site.Create()

	if rec.Error != nil {
		req.Serve(http.StatusInternalServerError, rec.Error, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec)
}

// @Title UpdateWebsite
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *ProfileController) Put() {
	body := &core.Profile{}
	key, err := req.GetKeyedRequest(body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = body.Update(key)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, nil)
}
