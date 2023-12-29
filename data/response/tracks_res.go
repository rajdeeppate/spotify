package response

type TracksRes struct {
	Id           int    `json:"id"`
	ISRC         string `json:"isrc"`
	Title        string `json:"title"`
	Artists      string `json:"artist"`
	SpotifyImage string `json:"sportifyimage"`
}
