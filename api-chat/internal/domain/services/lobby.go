package services

type Lobby struct {
	Users    map[*User]bool
	Messages chan[]byte
	Register chan *User

}

func NewLobby() *Lobby {
	return &Lobby{
		Messages: make(chan []byte),
	}
}

func (l *Lobby) Run() {
	for {
		select {
		case user := <- l.Register:
			l.Users[user] = true
		case message := <- l.Messages:
			for user := range l.Users {
				select {
				case user.Send <- message:
				default:
					close(user.Send)
					delete(l.Users, user)
				}
		}
		}
	}
}
