package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anudeep-mp/tracker/database"
	"github.com/anudeep-mp/tracker/helper"
	"github.com/anudeep-mp/tracker/model"
	"github.com/anudeep-mp/tracker/utilities"
)

func ServeHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to tracker api</h1>"))
}

func WatchStampHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	envionment := r.Header.Get("Environment")

	database.UpdateCollection(envionment)

	var watchStamp model.WatchStamp

	_ = json.NewDecoder(r.Body).Decode(&watchStamp)

	_, err := helper.PostWatchStamp(watchStamp)

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}
}

func GetWatchStampsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "GET")

	envionment := r.Header.Get("Environment")

	database.UpdateCollection(envionment)

	var users []model.UserStamp

	users, err := helper.GetWatchStamps()

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	result := model.GetWatchStampsResponse{
		UserCount: len(users),
		Users:     users,
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "Users fetched successfully", result)
}

func DeleteAllWatchStampsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	envionment := r.Header.Get("Environment")

	database.UpdateCollection(envionment)

	err := helper.DeleteAllWatchStamps()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "All users deleted successfully", nil)
}
