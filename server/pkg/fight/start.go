package fight

import (
	"encoding/json"
	"log"
	"net/http"
)

type startReq struct {
	Payload []string `json:"payload"`
}

func Start(w http.ResponseWriter, r *http.Request) {
	var req startReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("could not unmarshal body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Printf("received: %v", req)
}
