package handles

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/core"
	"log"
	"net/http"
	"strings"
)

func ProfileColour(w http.ResponseWriter, r *http.Request) {
	prf := drx.FindParam(r, "profile")
	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	rec, err := core.GetDisplay(prf, claims["azp"].(string))

	if err != nil {
		log.Println("GetDisplay Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	content := rec.GetValue().(core.Content)
	colour := content.Colour.GenerateCSS()

	name := fmt.Sprintf("%s.css", prf)
	err = mix.Write(w, mix.Octet(name, strings.NewReader(colour)))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
