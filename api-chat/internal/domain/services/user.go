package services

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
)

type User struct {
	Lobby *Lobby
	Conn  *websocket.Conn
	Send  chan []byte
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func NewUser(lobby *Lobby, conn *websocket.Conn, send chan []byte) *User {
	return &User{
		Lobby: lobby,
		Conn:  conn,
		Send:  send,
	}
}

func (u *User) ReadMessages() {
	for {
		_, message, err := u.Conn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		u.Lobby.Messages <- message
	}
}

func (u *User) WriteMessages() {
	for {
		select {
		case message, ok := <- u.Send:
			if !ok {
				log.Print("error in read message to write")
			}
			err := u.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("error: %v", err)
			}
		}
	}
}