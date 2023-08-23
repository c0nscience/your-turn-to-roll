package session

type Session struct {
	Message []string `json:"message"`
}

var sessions = map[int]*Session{}

func SetMessage(id int, msg []string) {
	session, ok := sessions[id]
	if !ok {
		return
	}

	session.Message = msg
}
