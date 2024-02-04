package ws

import "github.com/gorilla/websocket"

type WebSocketClient struct {
	Conn *websocket.Conn
}
