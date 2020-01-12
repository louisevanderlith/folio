package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/folio/controllers/profile"
	"github.com/louisevanderlith/folio/controllers/theme"
	"github.com/louisevanderlith/folio/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()
	profiles := r.Group("/profile")
	profiles.POST("", profile.Create)
	profiles.GET("/:key", profile.View)

	r.GET("/profiles", profile.Get)
	r.GET("/profiles/:pagesize/*hash", profile.Search)

	r.GET("/theme/:name", theme.Get)

	err := r.Run(":8090")

	if err != nil {
		panic(err)
	}
}
