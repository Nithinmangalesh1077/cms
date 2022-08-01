package models

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/config"
	"github.com/jinzhu/gorm"
)

var ac *gorm.DB

type Episode struct {
	gorm.Model
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	ac = config.GetDB()
	ac.AutoMigrate(&Episode{})
}

func (e *Episode) CreateEpisode() *Episode {
	ac.NewRecord(e)
	ac.Create(&e)
	return e
}

func GetAllEpisodes() []Episode {
	var Episodes []Episode
	ac.Find(&Episodes)
	return Episodes
}

func GetEpisodeById(Id int64) (*Episode, *gorm.DB) {
	var getEpisode Episode
	ac := ac.Where("ID=?", Id).Find(&getEpisode)
	return &getEpisode, ac
}

func DeleteEpisode(ID int64) Episode {
	var episode Episode
	ac.Where("ID=?", ID).Delete(episode)
	return episode
}
