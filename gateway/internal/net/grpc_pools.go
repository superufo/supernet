package net

import (
	"errors"
	"github.com/supernet/gateway/pkg/log"
	"go.uber.org/zap"
	"google.golang.org/grpc/connectivity"
	"sync"
	"sync/atomic"

	"sort"
)

var (
	gmax = 10

	MaxClientPerGrpc = 3000

	ErrConnShutdown = errors.New("grpc conn shutdown")
)

type GrpcPools struct {
	SortGrpcClientConn

	mu sync.RWMutex

	// 客户端id = 远程socket ip端口+backurl 对应的多个hall grpc连接
	// 一个hall grpc连接可以被多个客户端复用
	clientToGrpc map[string]GrpcClientConn
	url          string
}

func NewGrpcPools(url string, grpcNum int) GrpcPools {
	log.ZapLog.With(zap.Any("url", url)).Info("NewGrpcPools")
	defer func() {
		if err := recover(); err != nil {
			log.ZapLog.With(zap.Any("err", err)).Error("NewGrpcPools")
		}
	}()

	p := GrpcPools{
		SortGrpcClientConn: make([]*GrpcClientConn, 0),
		mu:                 sync.RWMutex{},
		clientToGrpc:       make(map[string]GrpcClientConn, 100),
		url:                url,
	}

	for i := 0; i < grpcNum; i++ {
		c, _ := NewGrpcClientConn(url)

		c.Run()
		p.SortGrpcClientConn = append(p.SortGrpcClientConn, c)
		log.ZapLog.With(zap.Any("i", i)).Info("NewGrpcPools")
	}

	return p
}

func (p GrpcPools) AddBatchConnect() {
	freeNum := 0

	for i := 0; i < len(p.SortGrpcClientConn); i++ {
		conn := p.SortGrpcClientConn[i]
		if conn.linkClientNum == 0 {
			freeNum++
		}
	}

	if freeNum < 1 {
		for i := 0; i < 8; i++ {
			c, _ := NewGrpcClientConn(p.url)

			c.Run()
			p.SortGrpcClientConn = append(p.SortGrpcClientConn, c)
			log.ZapLog.With(zap.Any("i", i)).Info("NewGrpcPools")
		}
	}
}

func (p GrpcPools) Get(clientId string) (conn *GrpcClientConn, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	log.ZapLog.With(zap.String("clientId", clientId)).Info("Get")
	log.ZapLog.With(zap.Any("连接池数量", p.SortGrpcClientConn.Len())).Info("Get")
	// 如果client存在对应的grpc ,直接使用
	if c, ok := p.clientToGrpc[clientId]; ok {
		conn = &c

		if conn.clientConn == nil || p.checkState(conn) != nil {
			conn = nil
		} else {
			goto RET
		}
	}

	// 排序 直接取第一个
	sort.Sort(p.SortGrpcClientConn)

	for i := 0; i < len(p.SortGrpcClientConn); i++ {
		ln := atomic.LoadInt32(&p.SortGrpcClientConn[i].linkClientNum)
		// 取最小连接数目的
		for k := 0; k < gmax; k++ {
			if ln == int32(k) {
				conn = p.SortGrpcClientConn[i]

				//log.ZapLog.With(zap.Any("p.SortGrpcClientConn[i]................", i)).Info("Get")
				if conn != nil && p.checkState(conn) == nil && conn.linkClientNum < int32(MaxClientPerGrpc) {
					atomic.AddInt32(&conn.linkClientNum, 1)

					goto RET
				}
			}
		}
	}

	for k := 0; k < 10; k++ {
		conn, err = NewGrpcClientConn(p.url)
		if err == nil {
			conn.Run()
			goto RET
		}
	}

	//for i := 0; i < len(p.SortGrpcClientConn); i++ {
	//	ln := atomic.LoadInt32(&p.SortGrpcClientConn[i].linkClientNum)
	//	log.ZapLog.With(zap.Int32("ins.linkClientNum", ln)).Info("Get----")
	//}
RET:
	if err == nil {
		p.clientToGrpc[clientId] = *conn
	}

	return conn, err
}

func (p GrpcPools) Return(clientId string) error {
	// 排序
	p.mu.Lock()
	defer p.mu.Unlock()

	if c, ok := p.clientToGrpc[clientId]; ok {
		atomic.AddInt32(&c.linkClientNum, -1)
		delete(p.clientToGrpc, clientId)
		//sort.Sort(p.SortGrpcClientConn)
	}

	return nil
}

func (p GrpcPools) checkState(conn *GrpcClientConn) error {
	if conn == nil || conn.clientConn == nil {
		return errors.New("grpc conn nil")
	}

	state := conn.clientConn.GetState()
	switch state {
	case connectivity.Ready:
		return nil
	}

	return ErrConnShutdown
}

func (p GrpcPools) getState(conn *GrpcClientConn) connectivity.State {
	if conn != nil && conn.clientConn != nil {
		return conn.clientConn.GetState()
	} else {
		return -99
	}
}

// ActiveDeadLink 激活所有失去的连接的 grpc
// 当连接不为激活状态的时候,重新生成一个新的连接
func (p GrpcPools) ActiveDeadLink(url string) {
	for _, gcc := range p.SortGrpcClientConn {
		//log.ZapLog.With(zap.Any("i:", i), zap.Any("state", p.getState(gcc))).Info("ActiveDeadLink.....")
		if p.checkState(gcc) != nil {
			//删除grpc链接
			gcc.Close()
			//p.SortGrpcClientConn = append(p.SortGrpcClientConn[:i], p.SortGrpcClientConn[i+1:]...)
			//log.ZapLog.With(zap.Any("url", url)).Info("ActiveDeadLink.....")

			p.mu.Lock()
			conn, err := NewGrpcClientConn(url)
			//log.ZapLog.With(zap.Any("err", err)).Info("ActiveDeadLink NewGrpcClientConn.....")
			if err != nil {
				conn.Run()
				p.SortGrpcClientConn = append(p.SortGrpcClientConn, conn)
			}
			p.mu.Unlock()
		}
	}
}

type SortGrpcClientConn []*GrpcClientConn

// Len 实现sort.Interface接口的获取元素数量方法
func (sgc SortGrpcClientConn) Len() int {
	return len(sgc)
}

// Less 实现sort.Interface接口的比较元素方法
func (sgc SortGrpcClientConn) Less(i, j int) bool {
	ci := atomic.LoadInt32(&sgc[i].linkClientNum)
	cj := atomic.LoadInt32(&sgc[j].linkClientNum)

	return ci < cj
}

// Swap 实现sort.Interface接口的交换元素方法
func (sgc SortGrpcClientConn) Swap(i, j int) {
	sgc[i], sgc[j] = sgc[j], sgc[i]
}
