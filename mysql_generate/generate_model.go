package mysql_generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

func GenModel(tables []*SQLTable, outFilePath string) error {
	for _, table := range tables {
		filePath := fmt.Sprintf("%s/%s.txt", outFilePath, table.Name)

		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		util.FilePrintf(f, definePO(table))
		util.FilePrintf(f, defineTableName(table))
		util.FilePrintf(f, defineDO(table))
		util.FilePrintf(f, definePOToDO(table))
		util.FilePrintf(f, defineDOToPO(table))

		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func definePO(table *SQLTable) (result string) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", util.HandlerFiledName(table.Name)))
	sb.WriteString("\tmysql_model.ColumnCreateModifyDeleteTime\n")
	for _, field := range table.Fields {
		if field.isBaseFiled {
			continue
		}
		sb.WriteString(fmt.Sprintf("\t%s\t%s\t`json:\"%s\"`\t// %s\n", util.HandlerFiledName(field.Name), field.Type, field.Name, field.Comment))
	}
	sb.WriteString("}\n\n")

	return sb.String()
}

func defineTableName(table *SQLTable) (result string) {
	result = fmt.Sprintf("func (p *%s) TableName() string {\n", util.HandlerFiledName(table.Name))
	result += fmt.Sprintf("\treturn \"%s\"\n}\n\n", table.Name)
	return result
}

func definePOToDO(table *SQLTable) (result string) {
	var sb strings.Builder

	poName := util.HandlerFiledName(table.Name)

	sb.WriteString(fmt.Sprintf("func POToDO%s(po *%s) (do *Do%s) {\n", poName, poName, poName))
	sb.WriteString(fmt.Sprintf("\tdo = &Do%s{\n", poName))
	for _, field := range table.Fields {
		fieldName := util.HandlerFiledName(field.Name)
		if fieldName == "Id" {
			sb.WriteString(fmt.Sprintf("\t\t%s:\tpo.ID,\n", fieldName))
			continue
		}
		sb.WriteString(fmt.Sprintf("\t\t%s:\tpo.%s,\n", fieldName, fieldName))
	}
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn do\n")
	sb.WriteString("}\n\n")
	return sb.String()
}
func defineDOToPO(table *SQLTable) (result string) {
	var sb strings.Builder
	var baseField = make([]*SQLField, 0, 5)

	poName := util.HandlerFiledName(table.Name)

	sb.WriteString(fmt.Sprintf("func DOToPO%s(do *Do%s) (po *%s) {\n", poName, poName, poName))
	sb.WriteString(fmt.Sprintf("\tpo = &%s{\n", poName))
	for _, field := range table.Fields {
		if field.isBaseFiled {
			baseField = append(baseField, field)
			continue
		}
		fieldName := util.HandlerFiledName(field.Name)
		sb.WriteString(fmt.Sprintf("\t\t%s:\tdo.%s,\n", fieldName, fieldName))
	}
	sb.WriteString("\t}\n")
	for _, field := range baseField {
		if field == nil {
			continue
		}
		fieldName := util.HandlerFiledName(field.Name)
		if fieldName == "Id" {
			sb.WriteString(fmt.Sprintf("\tpo.ID = do.%s\n", fieldName))
			continue
		}
		sb.WriteString(fmt.Sprintf("\tpo.%s = do.%s\n", fieldName, fieldName))
	}
	sb.WriteString("\treturn po\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func defineDO(table *SQLTable) (result string) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type Do%s struct {\n", util.HandlerFiledName(table.Name)))
	for _, field := range table.Fields {
		sb.WriteString(fmt.Sprintf("\t%s\t%s\t`json:\"%s\"`\t// %s\n", util.HandlerFiledName(field.Name), field.Type, field.Name, field.Comment))
	}
	sb.WriteString("}\n\n")
	return sb.String()
}
