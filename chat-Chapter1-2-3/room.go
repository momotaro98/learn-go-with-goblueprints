package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネル
	forward chan []byte
	// joinはチャットルームに参加しようとしているクライアントのためのチャネル
	join chan *client
	// leaveはチャットルームから退室しようとしているクライアントのためのチャネル
	leave chan *client
	// clientsには在室しているすべてのクライアントが保持される
	clients map[*client]bool

	// 複数のgoroutineがマップを同時に変更してもスレッドセーフにするためにjoinとleaveチャネルを利用する
}

// newRoomはすぐに利用できるチャットルームを生成して返す
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// runはjoin, leave, forwardを監視しそれぞれ処理をする
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// すべてのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

// ServeHTTPを実装するのでroomもHttpHandlerに渡せる
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil) // WebSocketによってHTTPというプロトコルはアップグレードした
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write() // クライアントへの表示は別スレッドで
	client.read()     // クライアントの書き込みを待つ
}
