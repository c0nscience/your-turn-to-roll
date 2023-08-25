package session

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type req struct {
	Characters []Character `json:"characters"`
}

func Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId, _ := strconv.ParseInt(vars["id"], 10, 32)

	if !Exists(int(sessionId)) {
		http.NotFound(w, r)
		return
	}

	var request req
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "could not unmarshal json", http.StatusInternalServerError)
		return
	}

	session, _ := Get(int(sessionId))
	session.Characters = request.Characters
}
