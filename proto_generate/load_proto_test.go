package proto_generate

import (
	"testing"
)

func TestLoadMessage(t *testing.T) {
	dataList, err := Load("../example.proto")
	if err != nil {
		return
	}
	// print dataList
	for _, data := range dataList {
		t.Logf("name : %#v,\t PBType: %v\n", data.Name, data.PBType)
		for _, filed := range data.Fields {
			t.Logf("Name:%v,\tType:%v,\tisSlice :%v,\tisBaseType:%v,\tComment :%s ,\t isRequired :%v \n", filed.Name, filed.Type, filed.IsSlice, filed.IsBaseType, filed.Comment, filed.IsRequired)
		}
		t.Log()
	}
}
