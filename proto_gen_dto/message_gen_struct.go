package proto_gen_dto

import (
	"fmt"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

type Translation struct {
	Content []string
}

func NewTranslation(content []string) *Translation {
	return &Translation{Content: content}
}

func (p *Translation) MessageToStruct() string {
	var sb strings.Builder
	structName := p.Content[0]
	sb.WriteString(p.defineStruct())
	sb.WriteString(p.definePBToDTO())
	sb.WriteString(p.definePBToDTOSlice(structName))
	sb.WriteString(p.defineDTOToPB())
	sb.WriteString(p.defineDTOToPBSlice(structName))
	return sb.String()
}

func (p *Translation) defineStruct() string {
	structName := p.Content[0]
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for i := 1; i+1 < len(p.Content); i = i + 2 {
		filedType := util.TypeConvert(p.Content[i])
		filedName := util.HandlerFiledName(p.Content[i+1])
		// 如果为message类型，则记录filedName与filedType的映射
		if _, ok := util.MessageMap[filedType]; ok {
			//fmt.Printf("filedName:%s,filedType :%s,messageMap : %#v\n", filedName, filedType, util.MessageMap)
			util.FiledMap[filedName] = filedType
		}
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(p.Content[i+1]))
		sb.WriteString("\t")
		sb.WriteString(util.TypeConvert(p.Content[i]))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("}\n\n"))
	return sb.String()
}

func (p *Translation) definePBToDTO() string {
	var sb strings.Builder
	structName := p.Content[0]

	sb.WriteString(fmt.Sprintf("func PBToDTO%s(param *pb.%s) (result *%s){\n", structName, structName, structName))
	sb.WriteString("\tif param == nil {\n")
	sb.WriteString("\t\t return nil\n")
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\tresult = &%s{\n", structName))
	for i := 1; i+1 < len(p.Content); i = i + 2 {
		sb.WriteString("\t\t")
		filedName := util.HandlerFiledName(p.Content[i+1])
		if filedType, ok := util.FiledMap[filedName]; ok {
			typeName := filedType
			typeName = strings.ReplaceAll(typeName, "*", "")
			typeName = strings.ReplaceAll(typeName, "[]", "")
			// 判断filedType是否为数组
			if strings.Index(filedType, "[]") > -1 {
				sb.WriteString(fmt.Sprintf("%s: PBToDTO%sSlice(param.%s),\n", filedName, typeName, filedName))
			} else {
				sb.WriteString(fmt.Sprintf("%s: PBToDTO%s(param.%s),\n", filedName, typeName, filedName))
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

func (p *Translation) definePBToDTOSlice(structName string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("func PBToDTO%sSlice(pbList []*pb.%s) (dtoList []*%s) {\n", structName, structName, structName))
	sb.WriteString(fmt.Sprintf("\tdtoList = make([]*%s,0,len(pbList))\n", structName))
	sb.WriteString("\t for _, item := range pbList {\n")
	sb.WriteString(fmt.Sprintf("\t\t dtoList = append(dtoList,PBToDTO%s(item))\n", structName))
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn dtoList\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func (p *Translation) defineDTOToPB() string {
	var sb strings.Builder
	structName := p.Content[0]

	sb.WriteString(fmt.Sprintf("func DTOToPB%s(param *%s) (result *pb.%s){\n", structName, structName, structName))
	sb.WriteString("\tif param == nil {\n")
	sb.WriteString("\t\t return nil\n")
	sb.WriteString("\t}\n")
	sb.WriteString(fmt.Sprintf("\tresult = &pb.%s{\n", structName))
	for i := 1; i+1 < len(p.Content); i = i + 2 {
		sb.WriteString("\t\t")
		filedName := util.HandlerFiledName(p.Content[i+1])
		if filedType, ok := util.FiledMap[filedName]; ok {
			typeName := filedType
			typeName = strings.ReplaceAll(typeName, "*", "")
			typeName = strings.ReplaceAll(typeName, "[]", "")
			// 判断filedType是否为数组
			if strings.Index(filedType, "[]") > -1 {
				sb.WriteString(fmt.Sprintf("%s: DTOToPB%sSlice(param.%s),\n", filedName, typeName, filedName))
			} else {
				sb.WriteString(fmt.Sprintf("%s: DTOToPB%s(param.%s),\n", filedName, typeName, filedName))
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

func (p *Translation) defineDTOToPBSlice(structName string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("func DTOToPB%sSlice(dtoList []*%s) (pbList []*pb.%s) {\n", structName, structName, structName))
	sb.WriteString(fmt.Sprintf("\tpbList = make([]*pb.%s,0,len(dtoList))\n", structName))
	sb.WriteString("\t for _, item := range dtoList {\n")
	sb.WriteString(fmt.Sprintf("\t\t pbList = append(pbList,DTOToPB%s(item))\n", structName))
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn pbList\n")
	sb.WriteString("}\n\n")
	return sb.String()
}
