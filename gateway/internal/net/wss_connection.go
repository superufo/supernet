package net

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/supernet/gateway/pkg/pack"
	"go.uber.org/zap"
	"sync"
	"time"

	"github.com/supernet/gateway/pkg/log"
)

// WsMessage 客户端读写消息
type WsMessage struct {
	messageType int
	data        []byte
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWaitSec = 180

	pongWait = pongWaitSec * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 1) / 18

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type WsConnection struct {
	wsConnect *websocket.Conn
	inChan    chan pack.WsMessage
	outChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex // 对closeChan关闭上锁
	isClosed bool       // 防止closeChan被关闭多次
}

func InitWsConnection(wsConn *websocket.Conn) (conn *WsConnection, err error) {
	conn = &WsConnection{
		wsConnect: wsConn,
		inChan:    make(chan pack.WsMessage, 0),
		outChan:   make(chan []byte, 0),
		closeChan: make(chan byte, 1),
	}

	return
}

func (conn *WsConnection) GetRemoteAddr() string {
	return conn.wsConnect.RemoteAddr().String()
}

func (conn *WsConnection) ReadMessage() (data pack.WsMessage, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("WsConnection is closeed")
	}
	return
}

func (conn *WsConnection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("WsConnection is closeed")
	}
	return err
}

func (conn *WsConnection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
func (conn *WsConnection) readLoop() {
	var (
		data []byte
		err  error
		t    int
	)

	//conn.wsConnect.SetReadLimit(maxMessageSize)
	//conn.wsConnect.SetReadDeadline(time.Now().Add(pongWait))
	//conn.wsConnect.SetPongHandler(func(string) error { conn.wsConnect.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		t, data, err = conn.wsConnect.ReadMessage()
		msg := pack.WsMessage{
			ContronMsg: t,
			Data:       data,
			Err:        err,
		}
		if err != nil {
			goto ERR
		}

		//Utils.Logger.Info(fmt.Sprintf("读取的数据为:%+v \n", data))
		//log.ZapLog.With(zap.Any("msg", msg)).Info("读取的数据")
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- msg:
		}
	}
ERR:
	conn.Close()
}

func (conn *WsConnection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
			//conn.wsConnect.SetWriteDeadline(time.Now().Add(writeWait))
			//发二进制数据
			if err = conn.wsConnect.WriteMessage(websocket.BinaryMessage, data); err != nil {
				goto ERR
			}
			log.ZapLog.With(zap.Any("data", data)).Info("reply cocos")
		}
	}

ERR:
	conn.Close()
}
