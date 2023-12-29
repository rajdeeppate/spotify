package service

import (
	"github.com/rajdeeppate/spotify.git/data/request"
	"github.com/rajdeeppate/spotify.git/data/response"
)

type TracksService interface {
	Create(tracks request.CreateTracksReq)
	Update(tracks request.UpdateTracksReq)
	Delete(tracksId int)
	FindById(tracksId int) response.TracksRes
	FindAll() []response.TracksRes
	FindByIsrc(isrc string) response.TracksRes
	FindByArtist(artist_name string) []response.TracksRes
}
