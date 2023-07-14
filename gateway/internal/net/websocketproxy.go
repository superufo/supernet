package net

//
//import (
//	"github.com/supernet/common/utils/bigendian"
//	"encoding/binary"
//	"fmt"
//	"github.com/supernet/gateway/pkg/log"
//	"sync"
//
//	"errors"
//	// "github.com/gorilla/websocket"
//
//	"github.com/supernet/gateway/config"
//	"github.com/supernet/gateway/pkg/gorilla/websocket"
//
//	"go.uber.org/zap"
//	"io"
//	"net"
//	"net/http"
//	"net/http/httputil"
//	"net/url"
//	"strconv"
//	"strings"
//	"time"
//)
//
//type Director func(incoming *http.Request, out http.Header)
//
//var (
//	DefaultUpgrader = &websocket.Upgrader{
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//		// 解决跨域问题
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}
//
//	//backUrlMaps = make(map[string]*url.URL)
//	pConf = config.NewProxyConfigs()
//)
//
//type ReverseProxy struct {
//	// DefaultDialer is a dialer with all fields set to the default zero values.
//	upgrader *websocket.Upgrader
//
//	backUrl *url.URL
//
//	dialer *websocket.Dialer
//}
//
//func NewReverseProxy() *ReverseProxy {
//	return &ReverseProxy{
//		upgrader: DefaultUpgrader,
//		backUrl:  nil,
//		dialer:   websocket.DefaultDialer,
//	}
//}
//
//func (rp *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
//	var (
//		mType uint16
//
//		p    []byte
//		wMsg int
//	)
//	// 默认
//	rp.backUrl = nil
//
//	rp.dialer = websocket.DefaultDialer
//	log.ZapLog.With().Info("ServeHTTP start....")
//
//	token := req.FormValue("token")
//	log.ZapLog.With(zap.Any("token", token)).Info("ServeHTTP....")
//
//	// 非 websocket 走http 代理
//	var once sync.Once
//	once.Do(func() {
//		pConf.GetAllMsgStrategy() // 解析消息对应的Strategy
//	})
//	pConf.GetMsgUrls(req) // 获取所有的消息对应的url
//	//log.ZapLog.With(zap.Any("urls", pConf.urls)).Info("ServeHTTP")
//	//log.ZapLog.With(zap.Any("strategys", pConf.strategys)).Info("ServeHTTP")
//
//	//log.ZapLog.With().Info("req.URL" + req.URL.String())
//	if websocket.IsWebSocketUpgrade(req) == false {
//		// 消息的协议号，通过http头msg来传递
//		msg := req.Header.Get("msg")
//		rp.backUrl = pConf.GetMsgUrlByStrategy(msg)
//
//		if rp.backUrl = pConf.GetMsgUrlByStrategy(msg); rp.backUrl != nil {
//			proxy := httputil.NewSingleHostReverseProxy(rp.backUrl)
//			proxy.ServeHTTP(rw, req)
//			log.ZapLog.With(zap.Any("backUrl", rp.backUrl)).Info("代理到后端的地址为")
//		} else {
//			//  没有找到对应的消息号码 本次请求中断
//			return
//		}
//	}
//
//	// 重新组装请求到后端的请求头信息
//	requestHeader := http.Header{}
//	if origin := req.Header.Get("Origin"); origin != "" {
//		requestHeader.Add("Origin", origin)
//	}
//	for _, prot := range req.Header[http.CanonicalHeaderKey("Sec-WebSocket-Protocol")] {
//		requestHeader.Add("Sec-WebSocket-Protocol", prot)
//	}
//	for _, cookie := range req.Header[http.CanonicalHeaderKey("Cookie")] {
//		requestHeader.Add("Cookie", cookie)
//	}
//	if req.Host != "" {
//		requestHeader.Set("Host", req.Host)
//	}
//
//	// 传递的返回头 代理端
//	upgradeHeader := http.Header{}
//	for _, hdr := range req.Header[http.CanonicalHeaderKey("Sec-WebSocket-Protocol")] {
//		upgradeHeader.Set("Sec-Websocket-Protocol", hdr)
//	}
//	for _, cookie := range req.Header[http.CanonicalHeaderKey("Cookie")] {
//		upgradeHeader.Set("Set-Cookie", cookie)
//	}
//
//	// 先通过握手协议升级
//	connPub, err := rp.upgrader.Upgrade(rw, req, upgradeHeader)
//	if err != nil {
//		log.ZapLog.With().Info(fmt.Sprintf("websocketproxy: couldn't upgrade %s", err))
//		return
//	}
//	defer connPub.Close()
//
//	//  当 TextMessage = 1  BinaryMessage = 2
//	if wMsg, p, err = connPub.ReadMessage(); err != nil {
//		log.ZapLog.With(zap.Any("err", err), zap.Any("wMsg", wMsg)).Error("请求body")
//	}
//	if binary.Size(p) > 2 {
//		mType = bigendian.FromUint16([2]byte{p[0], p[1]})
//	}
//
//	// websocket代理 如果发过来的 消息号等于后端配置的消息号码 代理到对应的websocket
//	rp.backUrl = pConf.GetMsgUrlByStrategy(strconv.Itoa(int(mType)))
//	if rp.backUrl != nil {
//		log.ZapLog.With(zap.Any("mType", mType), zap.Any("backUrl", rp.backUrl)).Info("代理到后端的地址为")
//
//		if strings.Contains(rp.backUrl.String(), "grpc") == true {
//
//		}
//	} else {
//		log.ZapLog.With().Info("没有找到对应的消息号码 本次请求中断")
//		return
//	}
//
//	// TODO: use RFC7239 http://tools.ietf.org/html/rfc7239
//	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
//		if prior, ok := req.Header["X-Forwarded-For"]; ok {
//			clientIP = strings.Join(prior, ", ") + ", " + clientIP
//		}
//		requestHeader.Set("X-Forwarded-For", clientIP)
//	}
//
//	requestHeader.Set("X-Forwarded-Proto", "http")
//	if req.TLS != nil {
//		requestHeader.Set("X-Forwarded-Proto", "https")
//	}
//
//	//  连接到后端 websocket start
//	connBackend, resp, err := rp.dialer.Dial(rp.backUrl.String(), requestHeader)
//	if err != nil {
//		log.ZapLog.With().Info(fmt.Sprintf("websocketproxy: couldn't dial to remote backend url %s", err))
//		if resp != nil {
//			// If the WebSocket handshake fails, ErrBadHandshake is returned
//			// along with a non-nil *http.Response so that callers can handle
//			// redirects, authentication, etcetera.
//			if err := copyResponse(rw, resp); err != nil {
//				log.ZapLog.With().Info(fmt.Sprintf("websocketproxy: couldn't write response after failed remote backend handshake: %s", err))
//			}
//		} else {
//			http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
//		}
//		return
//	}
//	defer connBackend.Close()
//	//  连接到后端 websocket  end
//
//	errClient := make(chan error, 1)
//	errBackend := make(chan error, 1)
//	replicateWebsocketConn := func(dst, src *websocket.Conn, errc chan error) {
//		for {
//			msgType, msg, err := src.ReadMessage()
//
//			if err != nil {
//				m := websocket.FormatCloseMessage(websocket.CloseNormalClosure, fmt.Sprintf("%v", err))
//				if e, ok := err.(*websocket.CloseError); ok {
//					if e.Code != websocket.CloseNoStatusReceived {
//						m = websocket.FormatCloseMessage(e.Code, e.Text)
//					}
//				}
//				errc <- err
//				dst.WriteMessage(websocket.CloseMessage, m)
//				break
//			}
//
//			if msgType == websocket.PingMessage || msgType == websocket.PongMessage || msgType == websocket.CloseMessage {
//				err = dst.WriteControl(msgType, msg, time.Now().Add(2*time.Second))
//			} else {
//				err = dst.WriteMessage(msgType, msg)
//			}
//
//			if err != nil {
//				errc <- err
//				break
//			}
//		}
//	}
//
//	go replicateWebsocketConn(connPub, connBackend, errClient)
//	go replicateWebsocketConn(connBackend, connPub, errBackend)
//
//	var message string
//	select {
//	case err = <-errClient:
//		message = "websocketproxy: Error when copying from backend to client: %v"
//	case err = <-errBackend:
//		message = "websocketproxy: Error when copying from client to backend: %v"
//	}
//	if e, ok := err.(*websocket.CloseError); !ok || e.Code == websocket.CloseAbnormalClosure {
//		log.ZapLog.With(zap.Any("error", err)).Info(message)
//	}
//}
//
//func (rp *ReverseProxy) pong() {
//
//}
//
//func copyHeader(dst, src http.Header) {
//	for k, vv := range src {
//		for _, v := range vv {
//			dst.Add(k, v)
//		}
//	}
//}
//
//func copyResponse(rw http.ResponseWriter, resp *http.Response) error {
//	copyHeader(rw.Header(), resp.Header)
//	rw.WriteHeader(resp.StatusCode)
//	defer resp.Body.Close()
//
//	_, err := io.Copy(rw, resp.Body)
//	return err
//}
//
//func PkgSPackage(protoNum uint16, protoData []byte) []byte {
//	var data []byte
//	if protoData != nil {
//		dataLen := binary.Size(protoData)
//		data = make([]byte, dataLen+2)
//	} else {
//		data = make([]byte, 2)
//	}
//
//	data[0] = byte(protoNum >> 8) // int8 == byte
//	data[1] = byte(protoNum)      //
//
//	copy(data[2:], protoData[:])
//
//	//protoNumByte := [2]byte{data[0], data[1]}
//	//log.ZapLog.With(zap.Any("protoNum", bigendian.FromUint16(protoNumByte))).Info("PkgSPackage打包")
//	return data
//}
//
//func UnPkgBgData(data []byte) (protoNum uint16, secret []byte, randNum []byte, protoData []byte, err error) {
//	byteNum := binary.Size(data)
//	if byteNum < 8 {
//		return 0, nil, nil, nil, errors.New(fmt.Sprintf("客户端包长度不够 %+v", data))
//	}
//
//	temp := [2]byte{data[0], data[1]}
//	protoNum = bigendian.FromUint16(temp)
//
//	secret = []byte{data[2], data[3]}
//	randNum = []byte{data[4], data[5], data[6], data[7]}
//
//	//  todo 长度计算怎么更优雅
//	protoBytes := make([]byte, byteNum-8)
//	for k, _ := range protoBytes {
//		protoBytes[k] = data[k+7]
//	}
//
//	return protoNum, secret, randNum, protoData, nil
//}
