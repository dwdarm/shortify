package main

import (
	"github.com/dwdarm/shortify/app/adapters"
	"github.com/dwdarm/shortify/app/domain/services"
	"github.com/dwdarm/shortify/app/main/handlers"
	persistences "github.com/dwdarm/shortify/app/persistences/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := adapters.NewPostgresSql()
	if err != nil {
		panic(err)
	}

	linkPersistence := persistences.NewLinkPostgresPersistence(db)
	linkService := services.NewLinkService(linkPersistence)
	linkHandler := handlers.NewLinkHandler(linkService)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	r.POST("api/links/", linkHandler.LinkCreate)
	r.GET("api/links/:slug", linkHandler.LinkGet)

	r.Run()
}
