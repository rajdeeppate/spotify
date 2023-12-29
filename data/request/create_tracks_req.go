package request

type CreateTracksReq struct {
	ISRC         string `validate:"required,max=200,min=1" json:"isrc"`
	Title        string `validate:"required,max=200,min=1" json:"title"`
	Artists      string `validate:"required,max=200,min=1" json:"artists"`
	SpotifyImage string `validate:"required,max=200,min=1" json:"sportifyimage"`
}
