package session

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"path/filepath"
)

type Session struct {
	Key        string            `json:"key"`
	Conns      []*websocket.Conn `json:"-"`
	Message    []string          `json:"-"`
	Characters []Character       `json:"characters"`
}

type Character struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Initiative float32 `json:"initiative"`
}

var sessions = map[int]*Session{}
var loaded = false

type MessageType string

const (
	UpdateMessage MessageType = "update"
)

const sessionStoreFile = "session.json"

type Message struct {
	Type    MessageType `json:"type"`
	Payload []string    `json:"payload"`
}

func Exists(id int) bool {
	_, ok := sessions[id]
	return ok
}

func Get(id int) (*Session, bool) {
	session, ok := sessions[id]
	return session, ok
}

func SetMessage(id int, msg []string) {
	session, ok := sessions[id]
	if !ok {
		return
	}

	session.Message = msg
}

func AddConnection(id int, conn *websocket.Conn) {
	session, ok := sessions[id]
	if !ok {
		return
	}

	session.Conns = append(session.Conns, conn)
}

func SendLastMessage(id int) {
	session, ok := sessions[id]
	if !ok {
		return
	}
	SendMessage(id, session.Message)
}
func SendMessage(id int, msg []string) {
	session, ok := sessions[id]
	if !ok {
		return
	}

	validConnections := []*websocket.Conn{}
	for _, conn := range session.Conns {
		err := conn.WriteJSON(Message{
			Type:    UpdateMessage,
			Payload: msg,
		})
		if err != nil {
			log.Printf("could not send message to client: %v\n", err)
			continue
		}

		validConnections = append(validConnections, conn)
	}
	session.Conns = validConnections
}

func Create(id int) *Session {
	res := &Session{
		Characters: []Character{},
	}
	sessions[id] = res
	return res
}

func Persist() {
	if !loaded {
		log.Printf("not sessions loaded yet ...")
		return
	}
	file, err := os.Create(sessionStoreFile)
	if err != nil {
		log.Printf("could not persist: %v", err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(sessions)
}

func Load() {
	loaded = true
	abs, _ := filepath.Abs(sessionStoreFile)
	log.Printf("session store file path: %s", abs)
	b, err := os.ReadFile(sessionStoreFile)
	if err != nil {
		log.Printf("could not load session store file: %v", err)
		return
	}

	_ = json.Unmarshal(b, &sessions)
}
