package fight

import (
	"encoding/json"
	"github.com/c0nscience/your-turn-to-roll/pkg/session"
	"log"
	"net/http"
)

type startReq struct {
	SessionId int      `json:"sessionId"`
	Payload   []string `json:"payload"`
}

func Start(w http.ResponseWriter, r *http.Request) {
	var req startReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("could not unmarshal body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Printf("received: %v", req)

	session.SetMessage(req.SessionId, req.Payload)
	session.SendMessage(req.SessionId, req.Payload)
	w.WriteHeader(http.StatusOK)
}
