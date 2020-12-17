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
	wordSlice := strings.Split(field, "_")
	var sb strings.Builder
	for _, word := range wordSlice {
		sb.WriteString(strings.Title(word))
	}
	return sb.String()
}

// 是否为Golang的基本类型
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

func TypeConvert(input string) string {
	switch input {
	case "double":
		return "float64"
	case "float":
		return "float32"
	}
	return input
}
