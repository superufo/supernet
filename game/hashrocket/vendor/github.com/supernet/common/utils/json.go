package utils

import (
	"bytes"
	"encoding/json"
)

//json.Marshal默认的SetEscapeHTML为true，会对'&','<','>'这些字符进行转义
func JSONMarshal(v interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return buf.Bytes(), err
}
