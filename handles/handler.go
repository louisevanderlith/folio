package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/kong/middle"
	"net/http"

	"github.com/gorilla/mux"
)

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "/", "Home", nil))
	m.AddItem(menu.NewItem("b", "/profiles", "Profiles", nil))
	m.AddItem(menu.NewItem("c", "/entities", "Entities", nil))
	m.AddItem(menu.NewItem("d", "/resources", "Resources", nil))
	m.AddItem(menu.NewItem("e", "/content", "Content Management", nil))

	return m
}

func SetupRoutes(clnt, scrt, securityUrl, managerUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, securityUrl, managerUrl, authorityUrl)
	r.HandleFunc("/callback", clntIns.Callback).Queries("state", "{state}", "token", "{token}").Methods(http.MethodGet)
	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), map[string]bool{"entity.info.search": true})).Methods(http.MethodGet)

	r.HandleFunc("/entities", clntIns.Middleware(GetEnitites(tmpl), map[string]bool{"entity.info.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchEntities(tmpl), map[string]bool{"entity.info.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchEntities(tmpl), map[string]bool{"entity.info.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/entities/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewEntity(tmpl), map[string]bool{"entity.info.view": true, "artifact.uploads.create": true})).Methods(http.MethodGet)

	r.HandleFunc("/content", clntIns.Middleware(GetAllContent(tmpl), map[string]bool{"cms.content.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchContent(tmpl), map[string]bool{"cms.content.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchContent(tmpl), map[string]bool{"cms.content.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/content/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewContent(tmpl), map[string]bool{"cms.content.view": true, "artifact.uploads.create": true, "entity.info.view": true})).Methods(http.MethodGet)

	r.HandleFunc("/profiles", clntIns.Middleware(GetProfiles(tmpl), map[string]bool{"secure.profile.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchProfiles(tmpl), map[string]bool{"secure.profile.search": true, "entity.info.view": true})).Methods(http.MethodGet)
	r.HandleFunc("/profiles/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewProfile(tmpl), map[string]bool{"secure.profile.view": true, "entity.info.view": true})).Methods(http.MethodGet)

	//crty.HandleFunc("/users", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, GetUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, SearchUsers(tmpl), map[string]bool{"secure.user.search": true})).Methods(http.MethodGet)
	//crty.HandleFunc("/users/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, authorityUrl, ViewUser(tmpl), map[string]bool{"secure.user.view": true})).Methods(http.MethodGet)

	r.HandleFunc("/resources", clntIns.Middleware(GetResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchResource(tmpl), map[string]bool{"secure.resource.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/resources/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewResource(tmpl), map[string]bool{"secure.resource.view": true})).Methods(http.MethodGet)

	return r
}
