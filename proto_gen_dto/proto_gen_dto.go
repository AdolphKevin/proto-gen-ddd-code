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
	rMessageContent := regexp.MustCompile("\\s+(\\w*)\\s+(\\w+)\\s*=")
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
		dtoToResponse(file, responseMatch[1], content)
	} else {
		util.FilePrintf(file, messageToStruct(content))
	}
}
func requestToDTO(file *os.File, prefixName string, content []string) {
	// generate dto struct
	structName := prefixName + "ReqDTO"
	util.FilePrintf(file, genRequestDTO(structName, content))

	// generate pb to dto func
	util.FilePrintf(file, genPBToDTO(prefixName, content))

}

func genRequestDTO(structName string, content []string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(content[i+1]))
		sb.WriteString("\t")
		sb.WriteString(util.TypeConvert(content[i]))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("}\n\n"))
	return sb.String()
}

func genPBToDTO(prefixName string, content []string) string {
	funcName := "PBToDTO" + prefixName
	structName := prefixName + "ReqDTO"
	pbStructName := prefixName + "Request"
	var sb strings.Builder
	// func PBToDTOBatchResultSendRecord(param *pb.BatchResultSendRecordRequest)(result *BatchResultSendRecordReqDTO) {
	sb.WriteString(fmt.Sprintf("func %s(param *pb.%s)(result *%s) {\n", funcName, pbStructName, structName))

	// new dto
	sb.WriteString(fmt.Sprintf("\tresult = &%s{\n", structName))
	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t\t")
		filedName := util.HandlerFiledName(content[i+1])
		sb.WriteString(fmt.Sprintf("%s: param.%s,\n", filedName, filedName))
	}
	sb.WriteString("\t}\n\n")
	// return dto
	sb.WriteString("\treturn result\n")

	sb.WriteString("}\n\n")
	return sb.String()
}

func dtoToResponse(file *os.File, prefixName string, content []string) {
	// generate response dto struct
	util.FilePrintf(file, genResponseDTO(prefixName, content))

	// generate dto to pb func
	util.FilePrintf(file, genDTOToPB(prefixName, content))
	return
}

func genResponseDTO(prefixName string, content []string) string {
	var sb strings.Builder
	structName := prefixName + "RespDTO"
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(content[i+1]))
		sb.WriteString("\t")
		sb.WriteString(util.TypeConvert(content[i]))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("}\n\n"))
	return sb.String()
}

func genDTOToPB(prefixName string, content []string) string {
	funcName := "DTOToPB" + prefixName
	structName := prefixName + "RespDTO"
	pbStructName := prefixName + "Response"

	var sb strings.Builder
	// func PBToDTOBatchResultSendRecord(param *pb.BatchResultSendRecordRequest)(result *BatchResultSendRecordReqDTO) {
	sb.WriteString(fmt.Sprintf("func %s(param *%s)(result *pb.%s) {\n", funcName, structName, pbStructName))

	// new pb result
	sb.WriteString(fmt.Sprintf("\tresult = &pb.%s{\n", pbStructName))
	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t\t")
		filedName := util.HandlerFiledName(content[i+1])
		sb.WriteString(fmt.Sprintf("%s: param.%s,\n", filedName, filedName))
	}
	sb.WriteString("\t}\n\n")
	// return dto
	sb.WriteString("\treturn result\n")

	sb.WriteString("}\n\n")
	return sb.String()
}
