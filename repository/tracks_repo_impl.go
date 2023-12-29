package repository

import (
	"errors"

	"github.com/rajdeeppate/spotify.git/data/request"
	"github.com/rajdeeppate/spotify.git/helper"
	"github.com/rajdeeppate/spotify.git/model"
	"gorm.io/gorm"
)

type TracksRepoImpl struct {
	Db *gorm.DB
}

func NewTracksRepoImpl(Db *gorm.DB) TracksRepo {
	return &TracksRepoImpl{Db: Db}
}

// Delete implements TracksRepo.
func (t *TracksRepoImpl) Delete(trackId int) {
	var track model.Tracks
	res := t.Db.Where("id = ?", trackId).Delete(&track)
	helper.ErrorPanic(res.Error)
}

// Save implements TracksRepo.
func (t *TracksRepoImpl) Save(tracks model.Tracks) {
	res := t.Db.Save(&tracks)
	helper.ErrorPanic(res.Error)
}

// FindALl implements TracksRepo.
func (t *TracksRepoImpl) FindALl() []model.Tracks {
	var tracks []model.Tracks
	res := t.Db.Find(&tracks)
	helper.ErrorPanic(res.Error)
	return tracks
}

// FindById implements TracksRepo.
func (t *TracksRepoImpl) FindById(trackId int) (tracks model.Tracks, err error) {
	var track model.Tracks
	res := t.Db.Find(&track, trackId)
	if res != nil {
		return track, nil
	} else {
		return track, errors.New("track not found")
	}
}

// Update implements TracksRepo.
func (t *TracksRepoImpl) Update(tracks model.Tracks) {
	var updateTrack = request.UpdateTracksReq{
		Id:    tracks.Id,
		Title: tracks.Title,
	}
	res := t.Db.Model(&tracks).Updates(updateTrack)
	helper.ErrorPanic(res.Error)

}

// FindByArtist implements TracksRepo.
func (t *TracksRepoImpl) FindByArtist(artist_name string) []model.Tracks {
	var tracks []model.Tracks
	res := t.Db.Where("artists LIKE ?", "%"+artist_name+"%").Find(&tracks)
	helper.ErrorPanic(res.Error)
	return tracks
}

// FindByIsrc implements TracksRepo.
func (t *TracksRepoImpl) FindByIsrc(isrc string) (tracks model.Tracks, err error) {
	var track model.Tracks
	res := t.Db.Where("isrc = ?", isrc).Find(&track)
	if res != nil {
		return track, nil
	} else {
		return track, errors.New("track not found")
	}
}
