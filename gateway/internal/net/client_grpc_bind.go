package net

// 调度 client 到对应的grpc连接池

type Bind struct {
	//cm     clientManager
	//gm     grpcPoolsManager

	Clients     map[string]*Client // 保持连接本grpc 的client
	grpcConnMap map[string]*GrpcClientConn
}
