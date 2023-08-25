package session

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Continue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId, _ := strconv.ParseInt(vars["id"], 10, 32)

	if !Exists(int(sessionId)) {
		http.NotFound(w, r)
		return
	}
	type resp struct {
		Id int `json:"id"`
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp{Id: int(sessionId)})
	if err != nil {
		http.Error(w, "could not marshal json", http.StatusInternalServerError)
		return
	}
}
