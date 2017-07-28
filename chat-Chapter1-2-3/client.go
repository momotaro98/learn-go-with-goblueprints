package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットをしている1人のユーザーを表す
type client struct {
	// socketはこのクライアントのためのWebSocket
	socket *websocket.Conn
	// sendはメッセージが送られるチャネル
	send chan []byte
	// roomはこのクライアントが参加しているチャットルーム
	room *room
}

// readはクライアントが書いたメッセージを読み込んでroomへそれを送る
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// writeはroomからのメッセージをクライアントへ送る
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
