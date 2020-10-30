package main

import (
	"flag"
	"github.com/louisevanderlith/folio/core"
	"github.com/louisevanderlith/folio/handles"
	"net/http"
	"time"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "folio", "Token target 'aud'")
	flag.Parse()

	db := core.CreateContext()
	defer db.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8090",
		Handler:      handles.SetupRoutes(*issuer, *audience),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
