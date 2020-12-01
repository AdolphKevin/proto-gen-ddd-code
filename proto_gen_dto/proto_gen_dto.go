package proto_gen_dto

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

func GenDTO(inFilePath, outFilePath string) (err error) {
	// message BatchResultSendRecordRequest {
	// ==> BatchResultSendRecordRequest、 {
	rMessage := regexp.MustCompile("message\\s+(\\w*)\\s+(.*)")
	//   string batch_no =
	// ==> string、 batch_no
	rMessageContent := regexp.MustCompile("\\s+(.*)\\s+(.*)\\s+=")
	rContentEnd := regexp.MustCompile("}")

	file, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var f *os.File
	f, err = os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(file)
	content := make([]string, 0)
	start := false
	for scanner.Scan() {
		messageMatch := rMessage.FindStringSubmatch(scanner.Text())
		messageContextMatch := rMessageContent.FindStringSubmatch(scanner.Text())
		endMatch := rContentEnd.FindStringSubmatch(scanner.Text())
		if len(messageMatch) > 0 {
			// 当遇到"{"时，开始将内容放入slice中
			//fmt.Println(messageMatch[1])
			content = append(content, messageMatch[1])
			// 标识开始
			start = true
		}
		if start == true && len(messageContextMatch) > 0 {
			//fmt.Println(messageContextMatch[1], messageContextMatch[2])
			content = append(content, messageContextMatch[1], messageContextMatch[2])
		}
		if start == true && len(endMatch) > 0 {
			//fmt.Println(endMatch[1])
			// 当遇到"}"时，开始处理数组中的内容，处理完后，清空slice
			handler(f, content)
			util.SliceClear(&content)
			// 标识结束
			start = false
		}
	}

	return
}

func handler(file *os.File, content []string) {

	rRequest := regexp.MustCompile("(.*)Request")
	rResponse := regexp.MustCompile("(.*)Response")

	requestMatch := rRequest.FindStringSubmatch(content[0])
	responseMatch := rResponse.FindStringSubmatch(content[0])

	if len(requestMatch) > 0 {
		requestToDTO(file, requestMatch[1], content)
	} else if len(responseMatch) > 0 {
		dtoToResponse(responseMatch[1], content)
	}
}
func requestToDTO(file *os.File, prefixName string, content []string) {
	structName := prefixName + "ReqDTO"
	var sb strings.Builder
	util.FilePrintf(file, fmt.Sprintf("type %s struct {\n", structName))
	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(content[i+1]))
		sb.WriteString("\t")
		sb.WriteString(util.TypeConvert(content[i]))
		sb.WriteString("\n")
	}
	util.FilePrintf(file, sb.String())
	util.FilePrintf(file, fmt.Sprintf("}\n\n"))
}
func dtoToResponse(prefixName string, content []string) {
	//structName := prefixName + "RespDTO"
	for i := 1; i < len(content); i++ {
	}
}
