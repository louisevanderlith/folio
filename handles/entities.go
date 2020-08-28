package handles

import (
	"github.com/louisevanderlith/folio/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/husk"
)

func GetEnitites(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Entities", tmpl, "./views/entities.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchEntities("A10")

		if err != nil {
			log.Println("Fetch Entities Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchEntities(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Entities", tmpl, "./views/entities.html")

	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchEntities(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewEntity(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Entity View", tmpl, "./views/entityview.html")

	return func(w http.ResponseWriter, r *http.Request) {

		key, err := husk.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchEntity(key.String())

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}
