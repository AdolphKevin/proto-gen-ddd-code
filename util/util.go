package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var MessageMap = make(map[string]struct{}) // is message type
var FiledMap = make(map[string]string)     // golang struct filedName mapping filedType

func FilePrintf(f *os.File, format string, a ...interface{}) {
	_, _ = io.WriteString(f, fmt.Sprintf(format, a...))
}

func SliceClear(s *[]string) {
	*s = (*s)[0:0]
}

// HandlerFiledName
// 1. remove charset _
// 2. initial to capitalize
func HandlerFiledName(field string) string {
	wordSlice := strings.Split(field, "_")
	var sb strings.Builder
	for _, word := range wordSlice {
		sb.WriteString(strings.Title(word))
	}
	return sb.String()
}

func TypeConvert(input string) (result string) {
	inputs := strings.Split(input, " ")
	if len(inputs) <= 1 {
		result = typeConvert(inputs[0])
		if !IsBaseType(inputs[0]) {
			result = "*" + result
			MessageMap[result] = struct{}{}
		}
		return result
	}

	// 意味着有repeated 关键字，需要转换成数组
	result = typeConvert(inputs[1])
	if IsBaseType(result) {
		result = "[]" + result
	} else {
		result = "[]*" + result
		MessageMap[result] = struct{}{}
	}
	return result
}

func IsBaseType(input string) bool {
	switch input {
	case "int32":
		return true
	case "int64":
		return true
	case "uint32":
		return true
	case "uint64":
		return true
	case "float":
		return true
	case "double":
		return true
	case "bool":
		return true
	case "string":
		return true
	}
	return false
}

func typeConvert(input string) string {
	switch input {
	case "double":
		return "float64"
	case "float":
		return "float32"
	}
	return input
}
