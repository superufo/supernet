package utils

import (
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

// StringsToInterfaces string数组转为interface数组
func StringsToInterfaces(src []string) []interface{} {
	res := make([]interface{}, 0, len(src))
	for i := 0; i < len(src); i++ {
		res = append(res, src[i])
	}
	return res
}

//Int64sToInterfaces int64数组转为interface数组
func Int64sToInterfaces(src []int64) []interface{} {
	res := make([]interface{}, 0, len(src))
	for i := 0; i < len(src); i++ {
		res = append(res, src[i])
	}
	return res
}

//UInt32sToInterfaces uint32数组转为interface数组
func UInt32sToInterfaces(src []uint32) []interface{} {
	res := make([]interface{}, 0, len(src))
	for i := 0; i < len(src); i++ {
		res = append(res, src[i])
	}
	return res
}

// set与slice转换

// SetToStrings set转string数组
func SetToStrings(set mapset.Set) []string {
	if set == nil {
		return nil
	}
	slice := make([]string, 0, set.Cardinality())
	it := set.Iterator()
	for vi := range it.C {
		if v, ok := vi.(string); ok {
			slice = append(slice, v)
		}
	}
	return slice
}

// SetToUInt32s set转uint32数组
func SetToUInt32s(set mapset.Set) []uint32 {
	if set == nil {
		return nil
	}
	slice := make([]uint32, 0, set.Cardinality())
	it := set.Iterator()
	for vi := range it.C {
		if v, ok := vi.(uint32); ok {
			slice = append(slice, v)
		}
	}
	return slice
}

// StringsToSet string数组转set
func StringsToSet(slice []string) mapset.Set {
	set := mapset.NewSet()
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

// TrimStringsToSet string数组转set，对string做trimSpace、去空值等处理
func TrimStringsToSet(slice []string) mapset.Set {
	set := mapset.NewSet()
	for _, v := range slice {
		v = strings.TrimSpace(v)
		if v != "" {
			set.Add(v)
		}
	}
	return set
}

// Uint32sToSet uint32数组转set
func Uint32sToSet(slice []uint32) mapset.Set {
	set := mapset.NewSet()
	for _, v := range slice {
		set.Add(v)
	}
	return set
}

func Byte2Rune(data []byte) []rune {
	var rs []rune
	for _, b := range data {
		rs = append(rs, rune(b))
	}
	return rs
}

func ToHexString(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}

func SliceCopy(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}
