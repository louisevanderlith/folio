package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/kong"
	"net/http"

	"github.com/gorilla/mux"
)

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "#home", "Home", nil))
	m.AddItem(menu.NewItem("b", "/profiles", "Profiles", nil))
	m.AddItem(menu.NewItem("c", "/entities", "Entities", nil))
	m.AddItem(menu.NewItem("d", "/resources", "Resources", nil))

	return m
}

func SetupRoutes(clnt, scrt, securityUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, Index(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)

	r.HandleFunc("/entities", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, GetEnitites(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchEntities(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchEntities(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, ViewEntity(tmpl), map[string]bool{"entity.info.view": true, "artifact.uploads.create": true})).Methods(http.MethodGet)

	r.HandleFunc("/profiles", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, GetProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, ViewProfile(tmpl), map[string]bool{"secure.profile.view": true})).Methods(http.MethodGet)

	//crty.HandleFunc("/users", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, GetUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, ViewUser(tmpl), map[string]bool{"secure.user.view": true})).Methods(http.MethodGet)

	r.HandleFunc("/resources", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, GetResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, ViewResource(tmpl), map[string]bool{"secure.resource.view": true})).Methods(http.MethodGet)

	return r
}
