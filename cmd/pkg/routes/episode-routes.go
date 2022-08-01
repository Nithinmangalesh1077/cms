package routes

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/controllers"
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterEpisodesRoutes = func(router *mux.Router) {

	router.HandleFunc("/episodes/", middleware.VerifyLogin(controllers.CreateEpisode)).Methods("POST")

	router.HandleFunc("/episodes/", controllers.GetEpisode).Methods("GET")

	router.HandleFunc("/episodes/{epdisodeId}", controllers.GetEpisodeById).Methods("GET")

	router.HandleFunc("/episodes/{epdisodeId}", middleware.VerifyLogin(controllers.UpdateEpisode)).Methods("PUT")

	router.HandleFunc("/episodes/{epdisodeId}", middleware.VerifyLogin(controllers.DeleteEpisode)).Methods("DELETE")

}
