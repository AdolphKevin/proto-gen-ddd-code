package proto_gen_dto

import (
	"fmt"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

func messageToStruct(content []string) string {
	var sb strings.Builder
	sb.WriteString(defineStruct(content))
	sb.WriteString(definePBToDTO(content))
	sb.WriteString(definePBToDTOSlice(content[0]))
	return sb.String()
}

func defineStruct(content []string) string {
	structName := content[0]
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for i := 1; i+1 < len(content); i = i + 2 {
		filedType := util.TypeConvert(content[i])
		filedName := util.HandlerFiledName(content[i+1])
		// 如果为message类型，则记录filedName与filedType的映射
		if _, ok := util.MessageMap[filedType]; ok {
			//fmt.Printf("filedName:%s,filedType :%s,messageMap : %#v\n", filedName, filedType, util.MessageMap)
			util.FiledMap[filedName] = filedType
		}
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(content[i+1]))
		sb.WriteString("\t")
		sb.WriteString(util.TypeConvert(content[i]))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("}\n\n"))
	return sb.String()
}

func definePBToDTO(content []string) string {
	var sb strings.Builder
	structName := content[0]

	sb.WriteString(fmt.Sprintf("func PBToDTO%s(param *pb.%s) (result *%s)\n", structName, structName, structName))
	sb.WriteString("\tif param == nil {\n")
	sb.WriteString("\t\t return nil\n")
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\tresult = &%s{\n", structName))
	for i := 1; i+1 < len(content); i = i + 2 {
		sb.WriteString("\t\t")
		filedName := util.HandlerFiledName(content[i+1])
		if filedType, ok := util.FiledMap[filedName]; ok {
			// 判断filedType是否为数组
			if strings.Index(filedType, "[]") > -1 {
				sb.WriteString(fmt.Sprintf("%s: PBToDTOSlice(param.%s),\n", filedName, filedName))
			} else {
				sb.WriteString(fmt.Sprintf("%s: PBToDTO(param.%s),\n", filedName, filedName))
			}
		} else {
			sb.WriteString(fmt.Sprintf("%s: param.%s,\n", filedName, filedName))
		}
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func definePBToDTOSlice(structName string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("func PBToDTO%sSlice(pbList []*pb.%s) (dtoList []*%s)) {\n", structName, structName, structName))
	sb.WriteString(fmt.Sprintf("\tdtoList = make([]*%s,0,len(pbList))\n", structName))
	sb.WriteString("\t for _, item := range pbList {\n")
	sb.WriteString(fmt.Sprintf("\t\t poList = append(poList,PBToDTO%s(item))\n", structName))
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn dtoList\n")
	sb.WriteString("}\n\n")
	return sb.String()
}
