package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/resources"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetResource(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Resources", tmpl, "./views/resources.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchResources("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		//result["Next"] = "resources/B10"
		//result["Previous"] = ""
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchResource(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Resources", tmpl, "./views/resources.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.AddMenu(FullMenu())
		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchResources(drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		//page, size := drx.GetPageData(r)
		//result["Next"] = fmt.Sprintf("%c%v", (page+1)+64, size)

		//if page != 1 {
		//	result["Previous"] = fmt.Sprintf("%c%v", (page-1)+64, size)
		//}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewResource(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Resource View", tmpl, "./views/resourceview.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchResource(key.String())

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
