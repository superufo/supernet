package iconv

import (
	"bytes"
	"testing"
)

const (
	NUM = 10000000
)

// TestStringConvert test
func TestStringConvert(t *testing.T) {
	origStr := "String convert test"
	convSlice := Str2BytesSlicePlus(origStr)
	byteSlice := []byte(origStr)

	if !bytes.Equal(convSlice, byteSlice) {
		t.Fail()
	}

	convSlice = Str2BytesSlicePlus2(origStr)
	if !bytes.Equal(convSlice, byteSlice) {
		t.Fail()
	}
}

// Benchmark_NormalConvert Run go test -bench="." -benchmem to verify efficiency
// go test -bench="." -benchmem
func Benchmark_NormalConvert(t *testing.B) {
	for count := 0; count < NUM; count++ {
		str := "This is string to byte slice convert test!!!"
		_ = normalString2BytesSlice(str)
	}
}

func Benchmark_PlusConvert(t *testing.B) {
	for count := 0; count < NUM; count++ {
		str := "This is string to byte slice convert test!!!"
		_ = Str2BytesSlicePlus(str)
	}
}

func Benchmark_PlusConvert2(t *testing.B) {
	for count := 0; count < NUM; count++ {
		str := "This is string to byte slice convert test!!!"
		_ = Str2BytesSlicePlus2(str)
	}
}
