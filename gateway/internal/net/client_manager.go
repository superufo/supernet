package net

import (
	"sync"
)

var (
	CManager *clientManager
)

func Init() {
	CManager = NewclientManager()
}

type clientManager struct {
	//Clients    map[*Client]bool
	Clients    map[string]*Client
	rwLock     sync.RWMutex
	IsClosed   bool
	Exit       chan interface{}
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client // 服务器主动关闭
}

func NewclientManager() *clientManager {
	exit := make(chan interface{}, 10000)
	broadcast := make(chan []byte, 10000)
	register := make(chan *Client, 10000)
	unregister := make(chan *Client, 10000)

	return &clientManager{
		Clients:    make(map[string]*Client),
		rwLock:     sync.RWMutex{},
		IsClosed:   false,
		Exit:       exit,
		Broadcast:  broadcast,
		Register:   register,
		Unregister: unregister,
	}
}

func (cm *clientManager) Run() {
	for {
		select {
		case client, ok := <-cm.Register:
			if ok {
				cm.Clients[client.sid] = client
				//log.ZapLog.With(zap.Any("Clients", cm.Clients[client.sid])).Info("clientManager")
			}
		case client, ok := <-cm.Unregister:
			if ok {
				delete(cm.Clients, client.sid)
			}
		case message, ok := <-cm.Broadcast:
			if ok {
				for _, client := range cm.Clients {
					client.Conn.WriteMessage(message)
				}
			}
		case <-cm.Exit:
			goto END
		}
	}
END:
}

func (cm *clientManager) Close() {
	cm.rwLock.Lock()
	defer cm.rwLock.Unlock()
	if !cm.IsClosed {
		cm.IsClosed = true
		close(cm.Broadcast)
		close(cm.Register)
		close(cm.Unregister)
		close(cm.Exit)
	}
}
