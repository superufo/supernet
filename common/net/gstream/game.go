package gstream

import "github.com/supernet/common/net/gstream/pb"

type Game interface {
	Run(quest <-chan *pb.StreamRequestData, response chan<- *pb.StreamResponseData)
}
