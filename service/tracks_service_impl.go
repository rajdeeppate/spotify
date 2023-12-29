package service

import (
	"fmt"

	"github.com/rajdeeppate/spotify.git/data/request"
	"github.com/rajdeeppate/spotify.git/data/response"
	"github.com/rajdeeppate/spotify.git/repository"

	"github.com/rajdeeppate/spotify.git/helper"
	"github.com/rajdeeppate/spotify.git/model"
)

type TracksServiceImpl struct {
	TracksRepository repository.TracksRepo
}

func NewTracksServiceImpl(trackRepository repository.TracksRepo) TracksService {
	return &TracksServiceImpl{
		TracksRepository: trackRepository,
	}
}

// Create implements TracksService
func (t *TracksServiceImpl) Create(tracks request.CreateTracksReq) {
	// err := t.Validate.Struct(tracks)
	// helper.ErrorPanic(err)
	trackModel := model.Tracks{
		Title:        tracks.Title,
		ISRC:         tracks.ISRC,
		Artists:      tracks.Artists,
		SpotifyImage: tracks.SpotifyImage,
	}
	t.TracksRepository.Save(trackModel)
}

// Delete implements TracksService
func (t *TracksServiceImpl) Delete(tracksId int) {
	t.TracksRepository.Delete(tracksId)
}

// FindAll implements TracksService
func (t *TracksServiceImpl) FindAll() []response.TracksRes {
	result := t.TracksRepository.FindALl()

	var tracks []response.TracksRes
	for _, value := range result {
		track := response.TracksRes{
			Id:           value.Id,
			Title:        value.Title,
			ISRC:         value.ISRC,
			Artists:      value.Artists,
			SpotifyImage: value.SpotifyImage,
		}
		tracks = append(tracks, track)
	}

	return tracks
}

// FindByArtist implements TracksService
func (t *TracksServiceImpl) FindByArtist(artist_name string) []response.TracksRes {
	result := t.TracksRepository.FindByArtist(artist_name)
	fmt.Println("assaasdd-------------", result)
	var tracks []response.TracksRes
	for _, value := range result {
		track := response.TracksRes{
			Id:           value.Id,
			Title:        value.Title,
			ISRC:         value.ISRC,
			Artists:      value.Artists,
			SpotifyImage: value.SpotifyImage,
		}
		tracks = append(tracks, track)
	}

	return tracks
}

// FindByIsrc implements TracksService
func (t *TracksServiceImpl) FindByIsrc(isrc string) response.TracksRes {
	trackData, err := t.TracksRepository.FindByIsrc(isrc)
	fmt.Println("-------------------", trackData)
	helper.ErrorPanic(err)

	trackResponse := response.TracksRes{
		Id:           trackData.Id,
		Title:        trackData.Title,
		ISRC:         trackData.ISRC,
		Artists:      trackData.Artists,
		SpotifyImage: trackData.SpotifyImage,
	}
	return trackResponse
}

// FindById implements TracksService
func (t *TracksServiceImpl) FindById(tracksId int) response.TracksRes {
	trackData, err := t.TracksRepository.FindById(tracksId)
	helper.ErrorPanic(err)

	trackResponse := response.TracksRes{
		Id:           trackData.Id,
		Title:        trackData.Title,
		ISRC:         trackData.ISRC,
		Artists:      trackData.Artists,
		SpotifyImage: trackData.SpotifyImage,
	}
	return trackResponse
}

// Update implements TracksService
func (t *TracksServiceImpl) Update(tracks request.UpdateTracksReq) {
	trackData, err := t.TracksRepository.FindById(tracks.Id)
	helper.ErrorPanic(err)
	trackData.Title = tracks.Title
	t.TracksRepository.Update(trackData)
}
