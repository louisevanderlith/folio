package main

import (
	"log"
	"os"
	"path"

	"github.com/louisevanderlith/folio/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/folio/core"
)

func main() {
	mode := os.Getenv("RUNMODE")
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	pubPath := path.Join(keyPath, pubName)

	core.CreateContext()
	defer core.Shutdown()

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(mode, appName, pubPath, enums.API)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv, host)
		beego.Run()
	}
}
