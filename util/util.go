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

func TypeConvert(input string) string {
	switch input {
	case "int":
		return "int32"
	case "int64":
		return "int64"
	}
	return input
}
