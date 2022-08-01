package models

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/config"
	"github.com/jinzhu/gorm"
)

var cd *gorm.DB

type Show struct {
	gorm.Model
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
}

func init() {
	config.Connect()
	cd = config.GetDB()
	cd.AutoMigrate(&Show{})
}

func (f *Show) CreateShow() *Show {
	cd.NewRecord(f)
	cd.Create(&f)
	return f
}

func GetAllShows() []Show {
	var Shows []Show
	cd.Find(&Shows)
	return Shows
}

func GetShowById(Id int64) (*Show, *gorm.DB) {
	var getShow Show
	cd := cd.Where("ID=?", Id).Find(&getShow)
	return &getShow, cd
}

func DeleteShow(ID int64) Show {
	var show Show
	cd.Where("ID=?", ID).Delete(show)
	return show
}
