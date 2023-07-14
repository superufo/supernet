package utils

import (
	"errors"
	"fmt"
	"reflect"
)

// StructCopy 结构体字段值复制
// from和to都必须传结构体对象指针
func StructCopy(from interface{}, to interface{}) error {
	var (
		fromValue = indirect(reflect.ValueOf(from))
		toValue   = indirect(reflect.ValueOf(to))
		fromType  = indirectType(fromValue.Type())
		toType    = indirectType(toValue.Type())
	)

	if !fromValue.CanAddr() || !toValue.CanAddr() {
		return errors.New("must be addressable")
	}

	// 判断类型
	if fromType.Kind() != reflect.Struct || toType.Kind() != reflect.Struct {
		return errors.New("must be struct")
	}

	//取出from的结构体字段
	fromTypeFields := deepFields(fromType)

	for _, field := range fromTypeFields {
		name := field.Name
		src := fromValue.FieldByName(name)
		dest := toValue.FieldByName(name)

		// 如果dest中不存在该字段，直接跳过
		if !dest.IsValid() {
			continue
		}

		if src.IsValid() {
			// from和to中字段名和类型一致时，才进行复制。否则报错
			if src.Kind() != dest.Kind() {
				srcStructName := reflect.TypeOf(from).String()
				destStructName := reflect.TypeOf(to).String()
				return fmt.Errorf("The Kind of Field %s is different in Struct %s and %s ",
					name, srcStructName, destStructName)
			} else {
				if dest.CanSet() {
					set(dest, src)
				}
			}
		}
	}

	return nil
}

// indirect 检查Value，指针需要用Elem()
func indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

// indirectType 检查Type，指针需要用Elem()
func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType
}

// deepFields 取出结构体字段
func deepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	if reflectType = indirectType(reflectType); reflectType.Kind() == reflect.Struct {
		for i := 0; i < reflectType.NumField(); i++ {
			v := reflectType.Field(i)
			if v.Anonymous {
				// 嵌套
				fields = append(fields, deepFields(v.Type)...)
			} else {
				fields = append(fields, v)
			}
		}
	}

	return fields
}

// set 反射Value的字段赋值
func set(to, from reflect.Value) bool {
	if from.IsValid() {
		if to.Kind() == reflect.Ptr {
			//set `to` to nil if from is nil
			if from.Kind() == reflect.Ptr && from.IsNil() {
				to.Set(reflect.Zero(to.Type()))
				return true
			} else if to.IsNil() {
				to.Set(reflect.New(to.Type().Elem()))
			}
			to = to.Elem()
		}

		if from.Type().ConvertibleTo(to.Type()) {
			to.Set(from.Convert(to.Type()))
		} else if from.Kind() == reflect.Ptr {
			return set(to, from.Elem())
		} else {
			return false
		}
	}
	return true
}
