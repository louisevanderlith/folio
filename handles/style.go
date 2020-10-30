package handles

import (
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/core"
	"log"
	"net/http"
	"strings"
)

func ProfileColour(w http.ResponseWriter, r *http.Request) {
	realm := drx.FindParam(r, "realm")
	client := drx.FindParam(r, "client")
	rec, err := core.Context().GetDisplay(realm, client)

	if err != nil {
		log.Println("GetDisplay Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	content := rec.GetValue().(core.Content)
	colour := content.Colour.GenerateCSS()

	name := fmt.Sprintf("%s.css", realm)
	err = mix.Write(w, mix.Octet(name, strings.NewReader(colour)))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
