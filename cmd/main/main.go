package main

import (
	"log"
	"net/http"

	"github.com/Nithinmangalesh1077/cms/cmd/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	routes.RegisterEpisodesRoutes(r)
	routes.RegisterSeasonRoutes(r)
	routes.RegisterShowRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
