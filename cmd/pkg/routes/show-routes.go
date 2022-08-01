package routes

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/controllers"
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterShowRoutes = func(router *mux.Router) {

	router.HandleFunc("/show/", middleware.VerifyLogin(controllers.CreateShow)).Methods("POST")

	router.HandleFunc("/show/", controllers.GetShow).Methods("GET")

	router.HandleFunc("/show/{showId}", controllers.GetShowById).Methods("GET")

	router.HandleFunc("/show/{showId}", middleware.VerifyLogin(controllers.UpdateShow)).Methods("PUT")

	router.HandleFunc("/show/{showId}", middleware.VerifyLogin(controllers.DeleteShow)).Methods("DELETE")

}
