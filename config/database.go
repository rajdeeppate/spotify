package config

import (
	"fmt"

	"github.com/rajdeeppate/spotify.git/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host          = "127.0.0.1"
	port          = 3306
	user          = "admin"
	password      = "Medika@2022"
	dbName        = "gocrud"
	ClientID      = "aee0e13ebbec4e73a3c5f646e63e5edd"
	ClientSecret  = "75e52caebe0746d584e6b46b305099a3"
	SpotifyAPIURL = "https://api.spotify.com/v1"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	helper.ErrorPanic(err)
	return db
}
