package bigendian

import (
	"encoding/binary"
	"unsafe"
)

// IsLittleEndian 判断是否是小端
func IsLittleEndian() bool {
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01
	// 大端(16进制)：00 00 00 01
	// 小端(16进制)：01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

// ToUint64 helper function which converts a uint64 to a []byte in Big Endian
func ToUint64(n uint64) [8]byte {
	s := make([]byte, 8)
	binary.BigEndian.PutUint64(s, n)
	a := [8]byte{}
	copy(a[:], s[:8])
	return a
}

// FromUint64 helper function which converts a big endian []byte to a uint64
func FromUint64(data [8]byte) uint64 {
	ui64 := binary.BigEndian.Uint64(data[:])
	return ui64
}

// ToInt helper function which converts a int to a []byte in Big Endian
func ToInt(n int) [4]byte {
	s := make([]byte, 4)
	binary.BigEndian.PutUint32(s, uint32(n))
	a := [4]byte{}
	copy(a[:], s[:4])
	return a
}

// FromInt helper function which converts a big endian []byte to an int
func FromInt(data [4]byte) int {
	ui32 := binary.BigEndian.Uint32(data[:])
	return int(ui32)
}

// ToInt16 helper function which converts a int to a []byte in Big Endian
func ToInt16(n uint16) [2]byte {
	s := make([]byte, 2)
	binary.BigEndian.PutUint16(s, uint16(n))
	a := [2]byte{}
	copy(a[:], s[:2])
	return a
}

// FromUint16 helper function which converts a big endian []byte to a uint64
func FromUint16(data [2]byte) uint16 {
	ui16 := binary.BigEndian.Uint16(data[:])
	return ui16
}
