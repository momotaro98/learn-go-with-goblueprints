package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/momotaro98/learn-go-with-goblueprints/trace"
	"github.com/stretchr/objx"
)

type room struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネル
	forward chan *message
	// joinはチャットルームに参加しようとしているクライアントのためのチャネル
	join chan *client
	// leaveはチャットルームから退室しようとしているクライアントのためのチャネル
	leave chan *client
	// clientsには在室しているすべてのクライアントが保持される
	clients map[*client]bool
	// avatarはアバターの情報を取得する
	avatar Avatar
	// tracerはチャットルーム上で行われた操作のログを受け取る
	tracer trace.Tracer

	// 複数のgoroutineがマップを同時に変更してもスレッドセーフにするためにjoinとleaveチャネルを利用する
}

// newRoomはすぐに利用できるチャットルームを生成して返す
func newRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		avatar:  avatar,
		tracer:  trace.Off(),
	}
}

// runはjoin, leave, forwardを監視しそれぞれ処理をする
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
			r.tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("クライアントが退出しました")
		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました: ", msg.Message)
			// すべてのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
					r.tracer.Trace(" -- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- 送信に失敗しました。クライアントをクリーンアップします")
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

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("クッキーの取得に失敗しました:", err)
		return
	}
	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write() // クライアントへの表示は別スレッドで
	client.read()     // クライアントの書き込みを待つ
}
