package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rajdeeppate/spotify.git/controller"
)

func NewRouter(tracksController *controller.TracksController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tracksRouter := baseRouter.Group("/tracks")
	tracksRouter.GET("", tracksController.FindAll)
	tracksRouter.GET("/:trackId", tracksController.FindById)
	tracksRouter.GET("/isrc/:isrc", tracksController.FindByIsrc)
	tracksRouter.GET("/artist/:artist_name", tracksController.FindByArtist)
	tracksRouter.POST("/createtrack", tracksController.Create)
	tracksRouter.PATCH("/:trackId", tracksController.Update)
	tracksRouter.DELETE("/:trackId", tracksController.Delete)

	return router
}
