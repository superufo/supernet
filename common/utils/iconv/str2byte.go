package iconv

import (
	"reflect"
	"unsafe"
)

func Byte2Str(bytes []byte) string {
	return *((*string)(unsafe.Pointer(&bytes)))
}

func Str2BytesSlicePlus(str string) []byte {
	bytesSlice := []byte{}                                                                                            //此处定义了一个空切片
	stringData := &(*(*reflect.StringHeader)(unsafe.Pointer(&str))).Data                                              //取得StringHeader的Data地址
	byteSliceData := &(*(*reflect.SliceHeader)(unsafe.Pointer(&bytesSlice))).Data                                     //取得SliceHeader的Data地址
	*byteSliceData = *stringData                                                                                      //将StringHeader.Data的值赋给SliceHeader.Data
	(*(*reflect.SliceHeader)(unsafe.Pointer(&bytesSlice))).Len = (*(*reflect.StringHeader)(unsafe.Pointer(&str))).Len //设置长度

	return bytesSlice
}

func Str2BytesSlicePlus2(str string) []byte {
	strSliceHeader := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	byteSlice := *(*[]byte)(unsafe.Pointer(&strSliceHeader))

	return byteSlice
}

// 效率低 使用上面两个函数有此函数100之效率 见下面 benchmem 函数，可测试
func normalString2BytesSlice(str string) []byte {
	return []byte(str)
}
