package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(issuer, audience string) http.Handler {
	r := mux.NewRouter()

	mw := open.BearerMiddleware(audience, issuer)

	//cnt := ins.Middleware("cms.content.view", scrt, DisplayContent)
	r.Handle("/display", mw.Handler(http.HandlerFunc(DisplayContent))).Methods(http.MethodGet)

	r.HandleFunc("/colour/{realm:[a-z]+}/{client:[a-z]+}", ProfileColour)

	//get := ins.Middleware("cms.content.search", scrt, GetContent)
	r.Handle("/content", mw.Handler(http.HandlerFunc(GetContent))).Methods(http.MethodGet)

	//view := ins.Middleware("cms.content.view", scrt, ViewContent)
	r.Handle("/content/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewContent))).Methods(http.MethodGet)

	//srch := ins.Middleware("cms.content.search", scrt, SearchContent)
	r.Handle("/content/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchContent))).Methods(http.MethodGet)
	r.Handle("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchContent))).Methods(http.MethodGet)

	//create := ins.Middleware("cms.content.create", scrt, CreateContent)
	r.Handle("/content", mw.Handler(http.HandlerFunc(CreateContent))).Methods(http.MethodPost)

	//update := ins.Middleware("cms.content.update", scrt, UpdateContent)
	r.Handle("/content/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(UpdateContent))).Methods(http.MethodPut)

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "cms.content.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
