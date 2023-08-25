package session

import (
	"github.com/gorilla/websocket"
	"log"
)

type Session struct {
	Conns      []*websocket.Conn
	Message    []string
	Characters []Character
}

type Character struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Initiative float32 `json:"initiative"`
}

var sessions = map[int]*Session{}

type MessageType string

const (
	UpdateMessage MessageType = "update"
)

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
