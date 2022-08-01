package routes

import (
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/controllers"
	"github.com/Nithinmangalesh1077/cms/cmd/pkg/middleware"
	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {

	router.HandleFunc("/movies/", middleware.VerifyLogin(controllers.CreateMovie)).Methods("POST")

	router.HandleFunc("/movies/", controllers.GetMovie).Methods("GET")

	router.HandleFunc("/movies/{movieId}", controllers.GetMovieById).Methods("GET")

	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.UpdateMovie)).Methods("PUT")

	router.HandleFunc("/movies/{movieId}", middleware.VerifyLogin(controllers.DeleteMovie)).Methods("DELETE")

	// router.HandleFunc("/movies/{name}", controllers.SearchMovie).Methods("GET")
	// // router.HandleFunc("/movies/", controllers.GetAllMoviessearch).Methods("GET")
	// router.HandleFunc("/movies/?limit,offset", controllers.GetMovieByPage).Methods("GET")

}
