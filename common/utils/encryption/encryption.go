//go:build encrypt
// +build encrypt

package encryption

// ByteEncrypt 加密
func ByteEncrypt(num int64, len int) []byte {
	x := num
	byteData := make([]byte, len)
	for i := 0; i < len; i++ {
		byteData[i] = byte(x & 255) //进行与运算
		x = x >> 8                  //右移8位
	}
	return byteData
}

// ByteDencrypt 解密
func ByteDencrypt(byteData []byte, len int) int32 {
	//测试位 []byte{48,191,8,0}
	//4-1-0  = 3  byteData[3] = 0
	//4-1-1  = 2  byteData[2] = 8
	//4-1-2  = 1  byteData[1] = 191
	//4-1-3  = 0  byteData[0] = 48
	//取最后一位
	x := int32(byteData[len-1])
	//len-1为第一个数没有进行位移，所以去掉循环
	for i := 0; i < len-1; i++ {
		index := len - 1 - i                //反向取数
		x = x<<8 + int32(byteData[index-1]) //保存上次计算结果
		//fmt.Println(x,i,index)
	}
	//fmt.Println(x)
	return x
}
