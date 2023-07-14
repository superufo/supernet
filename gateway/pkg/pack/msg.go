package pack

// WsMessage 客户端读写消息
type WsMessage struct {
	ContronMsg int
	Data       []byte
	Err        error
}
