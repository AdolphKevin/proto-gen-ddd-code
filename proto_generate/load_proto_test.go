package proto_generate

import (
	"fmt"
	"testing"
)

func TestLoadMessage(t *testing.T) {
	dataList, err := Load("../example.proto")
	if err != nil {
		return
	}
	// print dataList
	for _, data := range dataList {
		fmt.Printf("name : %#v,\t PBType: %v\n", data.Name, data.PBType)
		for _, filed := range data.Fields {
			fmt.Printf("Name:%v,\tType:%v,\tisSlice :%v,\tisBaseType:%v\n", filed.Name, filed.Type, filed.IsSlice, filed.IsBaseType)
		}
		fmt.Println()
	}
}
