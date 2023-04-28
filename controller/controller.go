package controller

import (
	"encoding/json"
	"net/http"

	"github.com/anudeep-mp/tracker/database"
	"github.com/anudeep-mp/tracker/helper"
	"github.com/anudeep-mp/tracker/model"
	"github.com/anudeep-mp/tracker/utilities"
	"github.com/gorilla/mux"
)

func ServeHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to tracker api</h1>"))
}

func WatchStampHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	environment := r.Header.Get("Environment")

	database.UpdateCollection(environment)

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

	environment := r.Header.Get("Environment")

	database.UpdateCollection(environment)

	var users []model.ResponseUserStamp

	users, err := helper.GetWatchStamps()

	if err != nil {
		utilities.ResponseWrapper(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	result := model.ResponseWatchStamp{
		UserCount: len(users),
		Users:     users,
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "Users fetched successfully", result)
}

func DeleteAllWatchStampsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	environment := r.Header.Get("Environment")

	database.UpdateCollection(environment)

	err := helper.DeleteAllWatchStamps()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "All users deleted successfully", nil)
}

func DeleteWatchStampHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	environment := r.Header.Get("Environment")

	database.UpdateCollection(environment)

	params := mux.Vars(r)

	err := helper.DeleteWatchStamp(params["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	utilities.ResponseWrapper(w, http.StatusOK, true, "User deleted successfully", nil)
}
