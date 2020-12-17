package mysql_generate

import (
	"bufio"
	"os"
	"regexp"
)

type SQLField struct {
	Name    string // 字段名
	Type    string // 字段类型
	Comment string // 字段描述
}

type SQLTable struct {
	Name   string      // 表名
	Fields []*SQLField // 字段
}

func Load(inFilePath string) (tables []*SQLTable, err error) {
	tables = make([]*SQLTable, 0)
	var table = &SQLTable{}

	file, err := os.Open(inFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	rTableName := regexp.MustCompile("(?i)CREATE TABLE.*\\.?`(\\w+)`")
	rField := regexp.MustCompile("`(?i)(\\w+)`\\s*(\\w+).*COMMENT\\s*'(.*)'")
	rEnd := regexp.MustCompile("(?i)\\)\\s*ENGINE")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tableNameMatch := rTableName.FindStringSubmatch(scanner.Text())
		fieldMatch := rField.FindStringSubmatch(scanner.Text())
		endMatch := rEnd.FindStringSubmatch(scanner.Text())

		if len(tableNameMatch) > 0 {

			table = &SQLTable{Name: tableNameMatch[1]}
			continue
		}
		if len(fieldMatch) > 0 {
			if isBaseField(fieldMatch[1]) {
				continue
			}

			field := &SQLField{
				Name:    fieldMatch[1],
				Type:    typeConvert(fieldMatch[2]),
				Comment: fieldMatch[3],
			}

			table.Fields = append(table.Fields, field)
			continue
		}
		if len(endMatch) > 0 {
			tables = append(tables, table)
			continue
		}
	}
	return
}

func typeConvert(input string) (output string) {
	switch input {
	case "varchar", "text":
		return "string"
	case "int", "bigint":
		return "int64"
	case "tinyint":
		return "int32"
	}
	return input
}

// 判断是否为MySQL表的基础字段
func isBaseField(field string) bool {
	switch field {
	case "id":
		return true
	case "create_time":
		return true
	case "modify_time":
		return true
	case "deleted_time":
		return true
	case "is_del":
		return true
	}
	return false
}
