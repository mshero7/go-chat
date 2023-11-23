package chat

import "github.com/gorilla/websocket"

type Client struct {
	id     string          // Client UUID
	socket *websocket.Conn // Client별 소켓
	send   chan []byte     // 메시지 받기 대기.
}

