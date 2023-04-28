package router

import (
	"github.com/anudeep-mp/tracker/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHomeHandler).Methods("GET")
	router.HandleFunc("/api/track", controller.WatchStampHandler).Methods("POST")
	router.HandleFunc("/api/watchstamps", controller.GetWatchStampsHandler).Methods("GET")
	router.HandleFunc("/api/watchstamps", controller.DeleteAllWatchStampsHandler).Methods("DELETE")
	router.HandleFunc("/api/watchstamp/{userId}", controller.DeleteWatchStampHandler).Methods("DELETE")

	return router
}
