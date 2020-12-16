package proto_generate

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

type PBField struct {
	Name       string
	Type       string
	IsSlice    bool
	IsBaseType bool
}

type PBData struct {
	Name   string
	PBType int
	Fields []*PBField
}

func Load(inFilePath string) (dataList []*PBData, dataMap map[string]*PBData, err error) {
	dataList = make([]*PBData, 0)
	dataMap = make(map[string]*PBData)
	var pbData = &PBData{}

	rMessage := regexp.MustCompile("message\\s+(\\w+)\\s*{")
	rMessageContent := regexp.MustCompile("\\s+(\\w*\\s*\\w+)\\s+(\\w+)\\s*=")
	rContentEnd := regexp.MustCompile("}")

	file, err := os.Open(inFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := false
	for scanner.Scan() {
		messageMatch := rMessage.FindStringSubmatch(scanner.Text())
		messageContextMatch := rMessageContent.FindStringSubmatch(scanner.Text())
		endMatch := rContentEnd.FindStringSubmatch(scanner.Text())
		if len(messageMatch) > 0 {
			// 当遇到带"{"message时，标记开始
			start = true
			// get message type
			messageType := getMessageType(messageMatch[1])
			pbData = &PBData{Name: messageMatch[1], PBType: messageType}

			continue
		}
		if start == true && len(messageContextMatch) > 0 {
			types := strings.Split(messageContextMatch[1], " ")
			if len(types) > 1 {
				isBaseType := util.IsBaseType(types[1])
				pbData.Fields = append(pbData.Fields, &PBField{Name: messageContextMatch[2], Type: types[1], IsSlice: true, IsBaseType: isBaseType})
			} else {
				isBaseType := util.IsBaseType(types[0])
				pbData.Fields = append(pbData.Fields, &PBField{Name: messageContextMatch[2], Type: types[0], IsBaseType: isBaseType})
			}
			continue
		}
		if start == true && len(endMatch) > 0 {
			// 遇到结束符，标识结束
			start = false

			// 将message添加到List
			dataList = append(dataList, pbData)
			// 并且建立message Name与字段的映射
			dataMap[pbData.Name] = pbData

			continue
		}
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
