package models

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/config"
	"github.com/jinzhu/gorm"
)

var ss *gorm.DB

type Season struct {
	gorm.Model

	Name string `json:"name"`

	Director string `json:"director"`

	Rating string `json:"rating"`

	Year string `json:"year"`
}

func init() {

	config.Connect()

	ss = config.GetDB()

	ss.AutoMigrate(&Season{})

}

func (m *Season) CreateSeason() *Season {

	ss.NewRecord(m)
	ss.Create(&m)

	return m

}

func GetAllSeasons() []Season {

	var Seasons []Season

	ss.Find(&Seasons)

	return Seasons

}

func GetSeasonById(Id int64) (*Season, *gorm.DB) {

	var getSeason Season

	ss := ss.Where("ID=?", Id).Find(&getSeason)

	return &getSeason, ss

}
func DeleteSeason(ID int64) Season {

	var season Season

	ss.Where("ID=?", ID).Delete(season)

	return season

}
