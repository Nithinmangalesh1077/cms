package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nithinmangalesh1077/cms/cmd/pkg/models"
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/utils"

	"github.com/gorilla/mux"
)

var NewSeason models.Season

func GetSeason(w http.ResponseWriter, r *http.Request) {
	newSeason := models.GetAllSeasons()
	res, _ := json.Marshal(newSeason)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSeasonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	seasonDetails, _ := models.GetSeasonById(ID)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSeason(w http.ResponseWriter, r *http.Request) {
	CreateSeason := &models.Season{}
	utils.ParseBody(r, CreateSeason)
	b := CreateSeason.CreateSeason()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	season := models.DeleteSeason(ID)
	res, _ := json.Marshal(season)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSeason(w http.ResponseWriter, r *http.Request) {
	var updateSeason = &models.Season{}
	utils.ParseBody(r, updateSeason)
	vars := mux.Vars(r)
	seasonId := vars["seasonId"]
	ID, err := strconv.ParseInt(seasonId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	seasonDetails, db := models.GetSeasonById(ID)
	if updateSeason.Name != "" {
		seasonDetails.Name = updateSeason.Name
	}
	if updateSeason.Director != "" {
		seasonDetails.Director = updateSeason.Director
	}
	if updateSeason.Rating != "" {
		seasonDetails.Rating = updateSeason.Rating
	}
	if updateSeason.Year != "" {
		seasonDetails.Year = updateSeason.Year
	}
	db.Save(&seasonDetails)
	res, _ := json.Marshal(seasonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
