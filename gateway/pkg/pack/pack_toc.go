package pack

import (
	"encoding/binary"
	bigendian "github.com/supernet/common/utils/bigendian"
)

// SPackage Pkg 发送给客户端包
// SPackage golang 编码网络字节序为小端
type SPackage struct {
	ProtoNum  uint16
	ProtoData []byte
}

func NewSPackage() SPackage {
	return SPackage{}
}

func (pack *SPackage) PkgSPackage(protoNum uint16, protoData []byte) []byte {
	var data []byte
	if protoData != nil {
		dataLen := binary.Size(protoData)
		data = make([]byte, dataLen+2)
	} else {
		data = make([]byte, 2)
	}

	data[0] = byte(protoNum >> 8) // int8 == byte
	data[1] = byte(protoNum)      //

	copy(data[2:], protoData[:])
	//protoNumByte := [2]byte{data[0], data[1]}
	//log.ZapLog.With(zap.Any("protoNum", bigendian.FromUint16(protoNumByte)), zap.Any("protoData", string(data[2:]))).Info("PkgSPackage打包")

	return data
}

func (pack *SPackage) UnPkgBgData(data []byte) {
	//temp := [2]byte{}
	//copy(temp[:], data[:2])
	temp := [2]byte{data[0], data[1]}
	pack.ProtoNum = bigendian.FromUint16(temp)

	//  todo 长度计算怎么更优雅
	dataLen := 0
	for i, _ := range data {
		dataLen = i + 1
	}

	protoBytes := make([]byte, dataLen-8)
	for k, _ := range protoBytes {
		protoBytes[k] = data[k+2]
	}

	pack.ProtoData = protoBytes
}
