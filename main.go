package main

import (
	"os"

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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("WEB_URL")},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("api/links", linkHandler.LinkCreate)
	r.GET("api/links/:slug", linkHandler.LinkGet)

	r.Run()
}
