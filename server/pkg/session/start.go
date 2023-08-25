package session

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(9999)
	sessions[id] = &Session{}
	type resp struct {
		Id int `json:"id"`
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp{Id: id})
	if err != nil {
		log.Printf("could not marshal json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
