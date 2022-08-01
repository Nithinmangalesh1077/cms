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

var NewShow models.Show

func GetShow(w http.ResponseWriter, r *http.Request) {
	newShow := models.GetAllShows()
	res, _ := json.Marshal(newShow)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetShowById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	showDetails, _ := models.GetShowById(ID)
	res, _ := json.Marshal(showDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateShow(w http.ResponseWriter, r *http.Request) {
	CreateShow := &models.Show{}
	utils.ParseBody(r, CreateShow)
	b := CreateShow.CreateShow()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	show := models.DeleteShow(ID)
	res, _ := json.Marshal(show)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateShow(w http.ResponseWriter, r *http.Request) {
	var updateShow = &models.Show{}
	utils.ParseBody(r, updateShow)
	vars := mux.Vars(r)
	showId := vars["showId"]
	ID, err := strconv.ParseInt(showId, 0, 0)
	if err != nil {
		fmt.Println("error while updating")

	}
	showDetails, cd := models.GetShowById(ID)
	if updateShow.Name != "" {
		showDetails.Name = updateShow.Name
	}
	if updateShow.Director != "" {
		showDetails.Director = updateShow.Director
	}
	if updateShow.Rating != "" {
		showDetails.Rating = updateShow.Rating
	}
	if updateShow.Year != "" {
		showDetails.Year = updateShow.Year
	}
	cd.Save(&showDetails)
	res, _ := json.Marshal(showDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
