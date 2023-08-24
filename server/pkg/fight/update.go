package fight

import (
	"encoding/json"
	"github.com/c0nscience/your-turn-to-roll/pkg/session"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type updateReq struct {
	Payload []string `json:"payload"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId, _ := strconv.ParseInt(vars["id"], 10, 32)

	var req updateReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session.SendMessage(int(sessionId), req.Payload)
	w.WriteHeader(http.StatusOK)
}
