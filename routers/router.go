package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/folio/controllers"

	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e *droxolite.Epoxy) {
	//Profile
	profCtrl := &controllers.ProfileController{}
	profGroup := droxolite.NewRouteGroup("profile", profCtrl)
	profGroup.AddRoute("Update Profile", "/", "PUT", roletype.Unknown, profCtrl.Put)
	profGroup.AddRoute("Profile by Name", "/{site:[a-zA-Z]+}", "GET", roletype.Unknown, profCtrl.GetOne)
	profGroup.AddRoute("All Profiles", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, profCtrl.Get)
	e.AddGroup(profGroup)

	//Theme
	themeCtrl := &controllers.ThemeController{}
	themeGroup := droxolite.NewRouteGroup("theme", themeCtrl)
	themeGroup.AddRoute("Profile Theme", "/{site:[a-zA-Z]+}", "GET", roletype.Unknown, themeCtrl.Get)
	e.AddGroup(themeGroup)

	/*ctrlmap := EnableFilters(s, host)

	profCtrl := controllers.NewProfileCtrl(ctrlmap)

	beego.Router("/v1/profile", profCtrl, "post:Post;put:Put")
	beego.Router("/v1/profile/:site", profCtrl, "get:GetOne")
	beego.Router("/v1/profile/all/:pagesize", profCtrl, "get:Get")*/
}

/*
func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.Owner
	emptyMap["PUT"] = roletype.Owner

	ctrlmap.Add("/v1/profile", emptyMap)

	beego.InsertFilter("/v1/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
*/
