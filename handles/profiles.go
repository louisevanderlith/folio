package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/resources"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/kong/prime"
	"html/template"
	"log"
	"net/http"
)

func GetProfiles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profiles", tmpl, "./views/profiles.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchProfiles("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		log.Println("Profiles:", result)
		//result["Next"] = "profiles/B10"
		//result["Previous"] = ""
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchProfiles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profiles", tmpl, "./views/profiles.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchProfiles(drx.FindParam(r, "pagesize"))

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

func ViewProfile(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profile View", tmpl, "./views/profileview.html")

	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Key Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result := make(map[string]interface{})
		src := resources.APIResource(http.DefaultClient, r)
		profile, err := src.FetchProfile(key.String())

		if err != nil {
			log.Println("Fetch Profile Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result["Profile"] = profile

		resources, err := src.FetchResources("A100")

		if err != nil {
			log.Println("Fetch Resources Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		result["Resources"] = CheckResources(profile.(prime.Profile), resources)

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func CheckResources(profile prime.Profile, resources records.Page) map[string]map[string]bool {
	result := make(map[string]map[string]bool)

	for _, client := range profile.Clients {
		result[client.Name] = make(map[string]bool)

		itor := resources.GetEnumerator()
		for itor.MoveNext() {
			rec := itor.Current().(hsk.Record)
			curr := rec.GetValue().(prime.Resource)
			result[client.Name][curr.Name] = false

			for _, resource := range client.AllowedResources {
				result[client.Name][resource] = curr.Name == resource
			}
		}
		itor.Reset()
	}

	return result
}
