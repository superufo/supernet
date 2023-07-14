package pack

import (
	"bytes"
	"encoding/binary"
	bigendian "github.com/supernet/common/utils/bigendian"
)

// Pkg 收到的客户端包
// Pkg golang 编码网络字节序为小端

type CPackage struct {
	ProtoNum  uint16
	ProtoData []byte
}

func NewCPackage() CPackage {
	return CPackage{}
}

func (pkg *CPackage) PkgBgData(protoNum uint16, secret [2]byte, randNum [4]byte, protoData []byte) []byte {
	var data []byte
	if protoData != nil {
		dataLen := binary.Size(protoData)
		data = make([]byte, dataLen+8)
	} else {
		data = make([]byte, 8)
	}

	data[0] = byte(protoNum)      // int8 == byte
	data[1] = byte(protoNum >> 8) //

	if protoData != nil {
		for k, _ := range protoData {
			data[k+8] = protoData[k]
		}
	}
	//copy(data[8:], protoData[:])

	return data
}

func (pkg *CPackage) UnPkgBgData(wmsg WsMessage) error {
	data := wmsg.Data // 多少byte
	r := bytes.NewReader(data)

	//if r.Size() < 8 {
	//	return errors.New(fmt.Sprintf("客户端包长度不够 %+v", data))
	//}
	t := readNByte(2, r)
	pkg.ProtoNum = bigendian.FromUint16([2]byte{t[0], t[1]})

	t = readNByte(0, r)
	pkg.ProtoData = readNByte(r.Len(), r)

	return nil
}

func readNByte(n int, r *bytes.Reader) (s []byte) {
	for i := 0; i < n; i++ {
		t, _ := r.ReadByte()
		s = append(s, t)
	}

	return s
}

// PkgSmBgData 小端发送
func (pkg *CPackage) PkgSmBgData(protoNum uint16, secret [2]byte, randNum [4]byte, protoData []byte) []byte {
	var data []byte
	if protoData != nil {
		dataLen := binary.Size(protoData)
		data = make([]byte, dataLen+8)
	} else {
		data = make([]byte, 8)
	}

	data[0] = byte(protoNum >> 8) // int8 == byte
	data[1] = byte(protoNum)      //

	if protoData != nil {
		for k, _ := range protoData {
			data[k+8] = protoData[k]
		}
	}

	return data
}
