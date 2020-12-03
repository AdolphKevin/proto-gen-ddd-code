package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

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
	fieldSlice := strings.Split(field, "_")
	var sb strings.Builder
	for _, word := range fieldSlice {
		sb.WriteString(strings.Title(word))
	}
	return sb.String()
}

func TypeConvert(input string) (result string) {
	inputs := strings.Split(input, " ")

	if len(inputs) <= 1 {
		if IsBaseType(inputs[0]) {
			result = typeConvert(inputs[0])
		} else {
			result = "*" + typeConvert(inputs[0])
		}
		return result
	}

	// 意味着有repeated 关键字，需要转换成数组
	result = typeConvert(inputs[1])
	if IsBaseType(result) {
		result = "[]" + result
	} else {
		result = "[]*" + result
	}
	return result
}

func typeConvert(input string) string {
	switch input {
	case "int":
		return "int32"
	case "int64":
		return "int64"
	}
	return input
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
	case "float32":
		return true
	case "float64":
		return true
	case "bool":
		return true
	case "string":
		return true
	}
	return false
}
