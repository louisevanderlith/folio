package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/folio/core"
)

type Profile struct {
}

func (req *Profile) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetProfiles(1, 10)

	return http.StatusOK, results
}

// @Title GetSites
// @Description Gets all sites
// @Success 200 {[]core.Profile} []core.Portfolio]
// @router /all/:pagesize [get]
func (req *Profile) Search(ctx context.Requester) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetProfiles(page, size)

	return http.StatusOK, results
}

// @Title GetSite
// @Description Gets customer website/profile
// @Param	site			path	string 	true		"customer website name OR ID"
// @Success 200 {core.Profile} core.Profile
// @router /:site [get]
func (req *Profile) View(ctx context.Requester) (int, interface{}) {
	siteParam := ctx.FindParam("site")

	key, err := husk.ParseKey(siteParam)

	if err != nil {
		byName, err := core.GetProfileByName(siteParam)

		if err != nil {
			return http.StatusNotFound, err
		}

		return http.StatusOK, byName
	}

	result, err := core.GetProfile(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, result
}

// @Title RegisterWebsite
// @Description Register a Website
// @Param	body		body 	folio.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Profile) Create(ctx context.Requester) (int, interface{}) {
	var site core.Profile
	err := ctx.Body(&site)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec := site.Create()

	if rec.Error != nil {
		return http.StatusInternalServerError, rec.Error
	}

	return http.StatusOK, rec
}

// @Title UpdateWebsite
// @Description Updates a Website
// @Param	body		body 	core.Profile	true		"body for service content"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Profile) Update(ctx context.Requester) (int, interface{}) {
	body := &core.Profile{}
	key, err := ctx.GetKeyedRequest(body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = body.Update(key)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, nil
}
