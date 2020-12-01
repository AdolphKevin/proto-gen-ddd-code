package main

import (
	"fmt"
	"github.com/AdolphKevin/proto-gen-ddd-code/proto_gen_dto"
	"github.com/AdolphKevin/proto-gen-ddd-code/proto_gen_server"
)

func main(){
	protoPath := "test.proto"
	genServerPath := "gen_result/server.txt"
	genDtoPath := "gen_result/dto.txt"

	err := proto_gen_server.GenServer(protoPath,genServerPath)
	if err != nil {
		return
	}

	err = proto_gen_dto.GenDTO(protoPath,genDtoPath)
	if err != nil {
		fmt.Println(err)
		return
	}

}
