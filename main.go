package main

import (
	"net/http"

	"github.com/rajdeeppate/spotify.git/config"
	"github.com/rajdeeppate/spotify.git/controller"
	"github.com/rajdeeppate/spotify.git/helper"
	"github.com/rajdeeppate/spotify.git/model"
	"github.com/rajdeeppate/spotify.git/repository"
	"github.com/rajdeeppate/spotify.git/router"
	"github.com/rajdeeppate/spotify.git/service"
)

func main() {

	// Database
	db := config.DatabaseConnection()

	db.Table("tracks").AutoMigrate(&model.Tracks{})

	// Repository
	tracksRepository := repository.NewTracksRepoImpl(db)

	// Service
	tracksService := service.NewTracksServiceImpl(tracksRepository)

	// Controller
	tracksController := controller.NewTracksController(tracksService)

	// Router
	routes := router.NewRouter(tracksController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
