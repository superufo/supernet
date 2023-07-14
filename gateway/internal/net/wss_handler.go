package net

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 5 * time.Second,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *WsConnection
	)

	// 完成ws协议的握手操作 Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = InitWsConnection(wsConn); err != nil {
		conn.Close()
	}

	client := NewClient(conn)
	client.CManager = CManager
	//client.ip = conn.GetRemoteAddr()

	// 客户端执行各种操作
	go client.Exec()
}
