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

	session, ok := Get(int(sessionId))
	if !ok {
		http.NotFound(w, r)
		return
	}
	type resp struct {
		Id         int         `json:"id"`
		Characters []Character `json:"characters"`
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp{
		Id:         int(sessionId),
		Characters: session.Characters,
	})
	if err != nil {
		http.Error(w, "could not marshal json", http.StatusInternalServerError)
		return
	}
}
