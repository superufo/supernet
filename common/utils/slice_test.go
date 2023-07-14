package utils

import (
	"fmt"
	"testing"
)

func TestSliceCopy(t *testing.T) {
	var a []byte = make([]byte, 0) //[]byte("dsfs")
	fmt.Println(SliceCopy(a))
}
