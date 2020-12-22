package proto_generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

var verifyMap = map[string]func(field string) string{}

func init() {
	verifyMap["interface"] = verifyInterface
	verifyMap["string"] = verifyString
	verifyMap["int64"] = verifyInt
	verifyMap["int32"] = verifyInt
	verifyMap["uint32"] = verifyInt
	verifyMap["uint64"] = verifyInt
	verifyMap["float64"] = verifyInt
	verifyMap["float32"] = verifyInt
}

func GenDTO(dataList []*PBMessage, outFilePath string) (err error) {
	f, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var dualWaySlice strings.Builder

	for _, data := range dataList {
		dualWaySlice.Reset()
		util.FilePrintf(f, DualWayMessage(data, dualWaySlice))
		util.FilePrintf(f, dualWaySlice.String())
	}
	return
}

// bilateral message and dto by PB file
func DualWayMessage(pbData *PBMessage, dualWaySlice strings.Builder) (result string) {
	var sb strings.Builder
	sb.WriteString(defineStruct(pbData))
	sb.WriteString(definePBToDTO(pbData, dualWaySlice))
	sb.WriteString(defineDTOToPB(pbData, dualWaySlice))
	return sb.String()
}

func defineStruct(pbData *PBMessage) (result string) {
	var sb strings.Builder
	switch pbData.PBType {
	case util.MessageType:
		sb.WriteString(fmt.Sprintf("type %s struct {\n", pbData.Name))
	case util.RequestType:
		sb.WriteString(fmt.Sprintf("type %sReqDTO struct {\n", pbData.Name))
	case util.ResponseType:
		sb.WriteString(fmt.Sprintf("type %sRespDTO struct {\n", pbData.Name))
	}

	for _, field := range pbData.Fields {
		sb.WriteString("\t")
		sb.WriteString(util.HandlerFiledName(field.Name))
		sb.WriteString("\t")
		if field.IsBaseType {
			sb.WriteString(field.Type)
		} else {
			if field.IsSlice {
				sb.WriteString("[]*" + field.Type)
			} else {
				sb.WriteString("*" + field.Type)
			}
		}
		sb.WriteString(fmt.Sprintf("\t `json:\"%s\"`", field.Name))
		sb.WriteString(fmt.Sprintf("\t // %s", field.Comment))
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("}\n\n"))
	return sb.String()
}

func definePBToDTO(pbData *PBMessage, dualWaySlice strings.Builder) (result string) {
	var sb strings.Builder
	var resultName = pbData.Name
	switch pbData.PBType {
	case util.RequestType:
		resultName = resultName + "ReqDTO"
	case util.ResponseType:
		return ""
	}
	sb.WriteString(fmt.Sprintf("func PBToDTO%s(param *pb.%s) (result *%s,err error){\n", pbData.Name, pbData.Name, resultName))
	sb.WriteString("\tif param == nil {\n")
	sb.WriteString("\t\t return nil,errors.New(\"param is nil\")\n")
	sb.WriteString("\t}\n")

	// required param valid
	for _, field := range pbData.Fields {
		if !field.IsRequired {
			continue
		}
		if field.IsBaseType {
			sb.WriteString(verifyMap[field.Type](field.Name))
		} else {
			sb.WriteString(verifyMap["interface"](field.Name))
		}
	}
	// pb to dto
	sb.WriteString(fmt.Sprintf("\tresult = &%s{\n", resultName))
	for _, field := range pbData.Fields {
		sb.WriteString("\t\t")
		fieldName := util.HandlerFiledName(field.Name)
		if field.IsBaseType {
			sb.WriteString(fmt.Sprintf("%s: param.%s,\n", fieldName, fieldName))
		} else {
			if field.IsSlice {
				dualWaySlice.WriteString(definePBToDTOSlice(field.Type))
				sb.WriteString(fmt.Sprintf("%s: PBToDTO%sSlice(param.%s),\n", fieldName, field.Type, fieldName))
			} else {
				sb.WriteString(fmt.Sprintf("%s: PBToDTO%s(param.%s),\n", fieldName, field.Type, fieldName))
			}
		}
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result,nil\n")
	sb.WriteString("}\n\n")

	return sb.String()
}

func defineDTOToPB(pbData *PBMessage, dualWaySlice strings.Builder) (result string) {
	var sb strings.Builder
	var paramName = pbData.Name
	switch pbData.PBType {
	case util.RequestType:
		return ""
	case util.ResponseType:
		paramName = paramName + "RespDTO"
	}

	sb.WriteString(fmt.Sprintf("func DTOToPB%s(param *%s) (result *pb.%s){\n", pbData.Name, paramName, pbData.Name))
	sb.WriteString("\tif param == nil {\n")
	sb.WriteString("\t\t return nil\n")
	sb.WriteString("\t}\n")

	sb.WriteString(fmt.Sprintf("\tresult = &pb.%s{\n", pbData.Name))
	for _, field := range pbData.Fields {
		sb.WriteString("\t\t")
		fieldName := util.HandlerFiledName(field.Name)
		if field.IsBaseType {
			sb.WriteString(fmt.Sprintf("%s: param.%s,\n", fieldName, fieldName))
		} else {
			if field.IsSlice {
				dualWaySlice.WriteString(defineDTOToPBSlice(field.Type))
				sb.WriteString(fmt.Sprintf("%s: DTOToPB%sSlice(param.%s),\n", fieldName, field.Type, fieldName))
			} else {
				sb.WriteString(fmt.Sprintf("%s: DTOToPB%s(param.%s),\n", fieldName, field.Type, fieldName))
			}
		}
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func definePBToDTOSlice(name string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("func PBToDTO%sSlice(pbList []*pb.%s) (dtoList []*%s) {\n", name, name, name))
	sb.WriteString(fmt.Sprintf("\tdtoList = make([]*%s,0,len(pbList))\n", name))
	sb.WriteString("\t for _, item := range pbList {\n")
	sb.WriteString(fmt.Sprintf("\t\t dtoList = append(dtoList,PBToDTO%s(item))\n", name))
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn dtoList\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func defineDTOToPBSlice(name string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("func DTOToPB%sSlice(dtoList []*%s) (pbList []*pb.%s) {\n", name, name, name))
	sb.WriteString(fmt.Sprintf("\tpbList = make([]*pb.%s,0,len(dtoList))\n", name))
	sb.WriteString("\t for _, item := range dtoList {\n")
	sb.WriteString(fmt.Sprintf("\t\t pbList = append(pbList,DTOToPB%s(item))\n", name))
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn pbList\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func verifyString(field string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\tif param.%s == \"\" { \n", field))
	sb.WriteString(fmt.Sprintf("\t\t return nil,errors.New(\"%s is empty\")\n", field))
	sb.WriteString("\t}\n")
	return sb.String()
}

func verifyInt(field string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\tif param.%s <= 0 { \n", field))
	sb.WriteString(fmt.Sprintf("\t\t return nil,errors.New(\"%s must large than 0\")\n", field))
	sb.WriteString("\t}\n")
	return sb.String()
}

func verifyInterface(field string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\tif param.%s == nil { \n", field))
	sb.WriteString(fmt.Sprintf("\t\t return nil,errors.New(\"%s is nil\")\n", field))
	sb.WriteString("\t}\n")
	return sb.String()
}
