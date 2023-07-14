package utils

import (
	"fmt"
	"testing"
)

//From form
type From struct {
	Name  string
	Age   int
	NumIn []int
	MIn   map[int]string
}

//To to
type To struct {
	Name  string
	Age   int
	Add   string
	NumIn []int
	MIn   map[int]string
}

//From2 from
type From2 struct {
	From
	Gender int
	Num    []int
	M      map[int]string
}

//To2 to2
type To2 struct {
	To
	Gender int
	School string
	Num    []int
	M      map[int]string
}

//TestStructCopy from to to
func TestStructCopy(t *testing.T) {
	var from = From{
		Name: "test",
		Age:  10,
	}

	var to = To{}
	err := StructCopy(&from, &to)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(to)
}

//TestStructCopy2 嵌套From2 转 To2
func TestStructCopy2(t *testing.T) {
	var from = From2{
		From: From{
			Name: "test",
			Age:  10,
		},
		Gender: 2,
	}

	var to = To{}
	err := StructCopy(&from, &to)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(to)
}

//TestStructCopy3 嵌套转嵌套
func TestStructCopy3(t *testing.T) {
	var from = From2{
		From: From{
			Name:  "test",
			Age:   10,
			NumIn: []int{1, 2, 3, 4},
			MIn:   map[int]string{1: "test"},
		},
		Gender: 2,
		Num:    []int{1, 2, 3, 4},
		M:      map[int]string{1: "test"},
	}

	var to = To2{}
	err := StructCopy(&from, &to)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(to)
}

// TestStructCopy4 From转嵌套
func TestStructCopy4(t *testing.T) {
	var from = From{
		Name: "test",
		Age:  10,
	}

	var to = To2{}
	err := StructCopy(&from, &to)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(to)
}
