package util

import (
	"reflect"
	"strings"
)

// 根据结构体字段获取 JSON 标签名称
func GetJSONName(obj interface{}, fieldName string) string {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	// 如果是指针，获取指向的类型
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	field, exists := typ.FieldByName(fieldName)
	if !exists {
		return fieldName
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldName
	}

	// 解析 JSON 标签，处理 "name,omitempty" 这种格式
	parts := strings.Split(jsonTag, ",")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}

	return fieldName
}

// 根据结构体字段获取 JSON 标签名称
// @deprecated 参考： GetJSONName
func GetJSONFieldName(structPtr interface{}, fieldName string) string {
	s := reflect.TypeOf(structPtr)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	field, found := s.FieldByName(fieldName)
	if !found {
		return fieldName // 如果找不到对应字段，返回原始字段名
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" || jsonTag == "-" {
		return fieldName
	}

	// 处理 json:"name,omitempty" 这样的标签
	jsonName := strings.Split(jsonTag, ",")[0]
	if jsonName == "" {
		return fieldName
	}

	return jsonName
}
