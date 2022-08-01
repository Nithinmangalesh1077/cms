package models

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/config"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "fmt"
)

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type Claims struct {
	Name string `json:"name"`

	Role string `json:"role"`

	jwt.StandardClaims
}

var db *gorm.DB

// type ProductModel struct {
// 	Db *mysql.DB
// }

type Movie struct {
	gorm.Model
	Name     string `json:"name"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
	Year     string `json:"year"`
	Poster   string `json:"poster"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var Movies []Movie
	db.Find(&Movies)
	return Movies
}

func GetMoviesById(Id int64) (*Movie, *gorm.DB) {
	var getMovie Movie
	db := db.Where("ID=?", Id).Find(&getMovie)
	return &getMovie, db
}

func DeleteMovie(ID int64) Movie {
	var movie Movie
	db.Where("ID=?", ID).Delete(movie)
	return movie
}

// func (productModel ProductModel) Contains(keyword string) ([]controllers.Movie, error) {
// 	rows, err := productModel.Db.Query("select * from movie where name like ?", "%"+keyword+"%")
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		movies := []controllers.Movie{}
// 		for rows.Next() {

// 			var name string
// 			var director string
// 			var rating string
// 			var year string
// 			var poster string
// 			err2 := rows.Scan(&name, &director, &rating, &year, &poster)
// 			if err2 != nil {
// 				return nil, err2
// 			} else {
// 				movie := controllers.Movie{name, director, rating, year, poster}
// 				movies = append(movies, movie)
// 			}
// 		}
// 		return movies, nil
// 	}
// }
// func SearchMovie(name string) (*Movie, *gorm.DB) {

// 	var namesearch Movie

// 	db := db.Where("NAME=?", name).Find(&namesearch)

// 	return &namesearch, db

// }
