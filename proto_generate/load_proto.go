package proto_generate

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

type PBMessage struct {
	Name   string
	PBType int        // PBType用于区分是为普通的message类型还是带了request/response尾缀的message
	Fields []*PBField // message中包含的字段
}

type PBField struct {
	Name       string // 字段名
	Type       string // 字段类型
	IsSlice    bool   // 字段是否为repeated数组
	IsBaseType bool   // 字段是否为go语言的基础类型
	Comment    string // PB字段的注释
	IsRequired bool   // 是否必填项
}

func Load(inFilePath string) (dataList []*PBMessage, err error) {
	dataList = make([]*PBMessage, 0)
	var pbData = &PBMessage{}
	var tempField *PBField = nil

	rMessage := regexp.MustCompile("message\\s+(\\w+)\\s*{")
	rMessageContent := regexp.MustCompile("\\s+(\\w*\\s*\\w+)\\s+(\\w+)\\s*=")
	rComment := regexp.MustCompile("//\\s*(.*)")
	rRequired := regexp.MustCompile("@required")
	rContentEnd := regexp.MustCompile("}")

	file, err := os.Open(inFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := false
	for scanner.Scan() {
		// ====================  read message start====================
		messageMatch := rMessage.FindStringSubmatch(scanner.Text())
		messageContextMatch := rMessageContent.FindStringSubmatch(scanner.Text())
		commentMatch := rComment.FindStringSubmatch(scanner.Text())
		requiredMatch := rRequired.FindStringSubmatch(scanner.Text())
		endMatch := rContentEnd.FindStringSubmatch(scanner.Text())
		if len(messageMatch) > 0 {
			// 当遇到带"{"message时，标记开始
			start = true
			// get message type
			messageType := getMessageType(messageMatch[1])
			pbData = &PBMessage{Name: messageMatch[1], PBType: messageType}

			continue
		}
		if start == true && len(requiredMatch) > 0 {
			if tempField == nil {
				tempField = &PBField{IsRequired: true}
			} else {
				tempField.IsRequired = true
			}
			continue
		}
		if start == true && len(commentMatch) > 0 {
			if tempField == nil {
				tempField = &PBField{Comment: commentMatch[1]}
			} else {
				tempField.Comment = commentMatch[1]
			}
			// 如果需要支持PB 字段末尾追加注释，eg :「 int32 test = 1 ; // 测试 」。则将此continue去掉。
			continue
		}
		if start == true && len(messageContextMatch) > 0 {
			field := &PBField{
				Name: messageContextMatch[2],
			}

			if tempField != nil {
				field.Comment = tempField.Comment
				field.IsRequired = tempField.IsRequired
			}
			// 获取类型，并判断是否为数组类型
			types := strings.Split(messageContextMatch[1], " ")
			if len(types) > 1 {
				field.IsSlice = true
				field.Type = util.TypeConvert(types[1])
			} else {
				field.Type = util.TypeConvert(types[0])
			}

			field.IsBaseType = util.IsBaseType(field.Type)
			pbData.Fields = append(pbData.Fields, field)

			tempField = nil
			continue
		}
		if start == true && len(endMatch) > 0 {
			// 遇到结束符，标识结束
			start = false

			// 将message添加到List
			dataList = append(dataList, pbData)

			continue
		}
		// ====================  read message end====================
	}
	return
}

func getMessageType(messageName string) (messageType int) {
	rRequest := regexp.MustCompile("(.*)Request")
	rResponse := regexp.MustCompile("(.*)Response")

	requestMatch := rRequest.FindStringSubmatch(messageName)
	responseMatch := rResponse.FindStringSubmatch(messageName)

	if len(requestMatch) > 0 {
		messageType = util.RequestType
	} else if len(responseMatch) > 0 {
		messageType = util.ResponseType
	} else {
		messageType = util.MessageType
	}
	return
}
