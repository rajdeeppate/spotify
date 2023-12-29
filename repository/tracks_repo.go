package repository

import "github.com/rajdeeppate/spotify.git/model"

type TracksRepo interface {
	Save(tracks model.Tracks)
	Update(tracks model.Tracks)
	Delete(trackId int)
	FindById(trackId int) (tracks model.Tracks, err error)
	FindALl() []model.Tracks
	FindByIsrc(isrc string) (tracks model.Tracks, err error)
	FindByArtist(artist_name string) []model.Tracks
}
