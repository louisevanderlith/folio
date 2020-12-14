package handles

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"
	"strings"
)

func GetContent(w http.ResponseWriter, r *http.Request) {
	results, err := core.Context().GetAllContent(1, 10)

	if err != nil {
		log.Println("Get Content Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func DisplayContent(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	iss := claims["iss"].(string)
	realm := iss[strings.LastIndex(iss, "/")+1:]
	clientId := claims["clientId"].(string)

	rec, err := core.Context().GetDisplay(realm, clientId)

	if err != nil {
		log.Println("Get Display Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec.GetValue()))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func ViewContent(w http.ResponseWriter, r *http.Request) {
	k := drx.FindParam(r, "key")
	key, err := keys.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.Context().GetContent(key)

	if err != nil {
		log.Println("Get Content Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(rec))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchContent(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	results, err := core.Context().GetAllContent(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func CreateContent(w http.ResponseWriter, r *http.Request) {
	var obj core.Content
	err := drx.JSONBody(r, &obj)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	k, err := core.Context().CreateContent(obj)

	if err != nil {
		log.Println("Create Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON(k))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := core.Content{}
	err = drx.JSONBody(r, &body)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.Context().UpdateContent(key, body)

	if err != nil {
		log.Println("Update Error", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
