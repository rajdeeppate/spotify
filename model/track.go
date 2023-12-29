package model

type Tracks struct {
	Id           int    `gorm:"type:int;prymary_key"`
	ISRC         string `gorm:"type:varchar(255)"`
	Title        string `gorm:"type:varchar(255)"`
	Artists      string `gorm:"type:varchar(255)"`
	SpotifyImage string `gorm:"type:text"`
}
