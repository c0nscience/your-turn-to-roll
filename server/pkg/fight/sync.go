package fight

import (
	"github.com/c0nscience/your-turn-to-roll/pkg/session"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)
import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Sync(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionId, _ := strconv.ParseInt(vars["id"], 10, 32)

	if !session.Exists(int(sessionId)) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session.AddConnection(int(sessionId), ws)
	session.SendLastMessage(int(sessionId))
}
