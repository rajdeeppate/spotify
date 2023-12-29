package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rajdeeppate/spotify.git/config"
	"github.com/rajdeeppate/spotify.git/data/request"
	"github.com/rajdeeppate/spotify.git/data/response"
	"github.com/rajdeeppate/spotify.git/helper"
	"github.com/rajdeeppate/spotify.git/service"
)

type TracksController struct {
	tracksService service.TracksService
}

func NewTracksController(service service.TracksService) *TracksController {
	return &TracksController{
		tracksService: service,
	}
}

// CreateTracks		godoc
// @Summary			Create tracks
// @Description		Save tracks data in Db.
// @Param			tracks body request.CreateTracksRequest true "Create tracks"
// @Produce			application/json
// @Tracks			tracks
// @Success			200 {object} response.Response{}
// @Router			/tracks [post]
func (controller *TracksController) Create(ctx *gin.Context) {
	fmt.Println("create tracks")

	createTracksReq := request.CreateTracksReq{}
	err := ctx.ShouldBindJSON(&createTracksReq)
	helper.ErrorPanic(err)

	trackData, err := getTrackInfo(config.ClientID, config.ClientSecret, config.SpotifyAPIURL, createTracksReq.ISRC)
	helper.ErrorPanic(err)

	controller.tracksService.Create(*trackData)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTracks		godoc
// @Summary			Update tracks
// @Description		Update tracks data.
// @Param			trackId path string true "update tracks by id"
// @Param			tracks body request.CreateTracksRequest true  "Update tracks"
// @Tracks			tracks
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tracks/{trackId} [patch]
func (controller *TracksController) Update(ctx *gin.Context) {
	fmt.Println("update tracks")
	updateTracksReq := request.UpdateTracksReq{}
	err := ctx.ShouldBindJSON(&updateTracksReq)
	helper.ErrorPanic(err)

	trackId := ctx.Param("trackId")
	id, err := strconv.Atoi(trackId)
	helper.ErrorPanic(err)
	updateTracksReq.Id = id

	controller.tracksService.Update(updateTracksReq)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTracks		godoc
// @Summary			Delete tracks
// @Description		Remove tracks data by id.
// @Produce			application/json
// @Tracks			tracks
// @Success			200 {object} response.Response{}
// @Router			/tracks/{trackID} [delete]
func (controller *TracksController) Delete(ctx *gin.Context) {
	fmt.Println("delete tracks")
	trackId := ctx.Param("trackId")
	id, err := strconv.Atoi(trackId)
	helper.ErrorPanic(err)
	controller.tracksService.Delete(id)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTracks 		godoc
// @Summary				Get Single tracks by id.
// @Param				trackId path string true "update tracks by id"
// @Description			Return the tahs whoes trackId valu mathes id.
// @Produce				application/json
// @Tracks				tracks
// @Success				200 {object} response.Response{}
// @Router				/tracks/{trackId} [get]
func (controller *TracksController) FindById(ctx *gin.Context) {
	fmt.Println("findbyid tracks")
	trackId := ctx.Param("trackId")
	id, err := strconv.Atoi(trackId)
	helper.ErrorPanic(err)

	trackRes := controller.tracksService.FindById(id)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   trackRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTracks 		godoc
// @Summary			Get All tracks.
// @Description		Return list of tracks.
// @Tracks			tracks
// @Success			200 {obejct} response.Response{}
// @Router			/tracks [get]
func (controller *TracksController) FindAll(ctx *gin.Context) {
	fmt.Println("findAll tracks")
	trackRes := controller.tracksService.FindAll()
	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   trackRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIsrc 		godoc
// @Summary			Get All tracks.
// @Description		Return list of tracks.
// @Tracks			tracks
// @Success			200 {obejct} response.Response{}
// @Router			/tracks [get]
func (controller *TracksController) FindByIsrc(ctx *gin.Context) {
	fmt.Println("FindByIsrc tracks")
	isrc := ctx.Param("isrc")

	trackRes := controller.tracksService.FindByIsrc(isrc)
	fmt.Println("----------------------", trackRes)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   trackRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByArtist 		godoc
// @Summary			Get All Artist.
// @Description		Return list of tracks.
// @Tracks			tracks
// @Success			200 {obejct} response.Response{}
// @Router			/tracks [get]
func (controller *TracksController) FindByArtist(ctx *gin.Context) {
	fmt.Println("FindByArtist tracks")
	artist_name := ctx.Param("artist_name")

	trackRes := controller.tracksService.FindByArtist(artist_name)

	webResponse := response.Res{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   trackRes,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func getTrackInfo(clientID, clientSecret, spotifyAPIURL, isrc string) (*request.CreateTracksReq, error) {
	// Retrieve Spotify access token
	token, err := getSpotifyAccessToken(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	// Fetch track info from Spotify API
	trackEndpoint := fmt.Sprintf("%s/tracks/%s", spotifyAPIURL, isrc)
	req, err := http.NewRequest("GET", trackEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("spotify api request failed with status code: %d", resp.StatusCode)
	}

	// Parse the response
	var spotifyTrack struct {
		Name    string
		Artists []struct {
			Name string
		}
		Album struct {
			Images []struct {
				URL string
			}
		}
		Popularity int `json:"popularity"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &spotifyTrack)
	if err != nil {
		return nil, err
	}

	// Create a Track struct with relevant information
	track := &request.CreateTracksReq{
		ISRC:         isrc,
		Title:        spotifyTrack.Name,
		Artists:      joinArtists(spotifyTrack.Artists),
		SpotifyImage: getBestImage(spotifyTrack.Album.Images),
	}

	return track, nil
}

func getSpotifyAccessToken(clientID, clientSecret string) (string, error) {
	authURL := "https://accounts.spotify.com/api/token"
	req, err := http.NewRequest("POST", authURL, strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("spotify api authentication failed with status code: %d", resp.StatusCode)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func joinArtists(artists []struct{ Name string }) string {
	var artistNames []string
	for _, artist := range artists {
		artistNames = append(artistNames, artist.Name)
	}
	return strings.Join(artistNames, ", ")
}

func getBestImage(images []struct{ URL string }) string {
	if len(images) > 0 {
		return images[0].URL
	}
	return ""
}
