package session

import (
	"encoding/json"
	"github.com/c0nscience/your-turn-to-roll/pkg/password"
	"math/rand"

	"net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(9999)

	idGenerationFailed := false
	tries := 10
	for Exists(id) {
		if tries <= 0 {
			idGenerationFailed = true
			break
		}
		id = rand.Intn(9999)
		tries--
	}

	if idGenerationFailed {
		http.Error(w, "could not generate id", http.StatusInternalServerError)
		return
	}

	type resp struct {
		Id  int    `json:"id"`
		Key string `json:"key"`
	}

	key := password.Generate(5)
	hash, err := password.Hash([]byte(key))
	if err != nil {
		http.Error(w, "generating password failed", http.StatusInternalServerError)
		return
	}

	newSession := Create(id)
	newSession.Key = string(hash)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp{
		Id:  id,
		Key: key,
	})
	if err != nil {
		http.Error(w, "could not marshal json", http.StatusInternalServerError)
		return
	}

}
