package request

type UpdateTracksReq struct {
	Id           int    `validate:"required"`
	ISRC         string `validate:"required,max=200,min=1" json:"isrc"`
	Title        string `validate:"required,max=200,min=1" json:"title"`
	Artists      string `validate:"required,max=200,min=1" json:"artist"`
	SpotifyImage string `validate:"required,max=200,min=1" json:"sportifyimage"`
}
