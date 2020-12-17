package main

import (
	"fmt"

	"github.com/AdolphKevin/proto-gen-ddd-code/mysql_generate"

	"github.com/AdolphKevin/proto-gen-ddd-code/proto_generate"

	"github.com/AdolphKevin/proto-gen-ddd-code/proto_gen_server"
)

func main() {
	protoPath := "example.proto"
	genServerPath := "gen_result/server.txt"
	genDtoPath := "gen_result/dto.txt"

	createTablePath := "./create_table.sql"
	genModelPath := "gen_result/mysql_model"

	err := proto_gen_server.GenServer(protoPath, genServerPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取 proto 文件
	dataList, err := proto_generate.Load(protoPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 生成dto
	err = proto_generate.GenDTO(dataList, genDtoPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 读取create table的建表SQL
	tables, err := mysql_generate.Load(createTablePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 生成model
	err = mysql_generate.GenModel(tables, genModelPath)
	if err != nil {
		fmt.Println(err)
		return
	}

}
