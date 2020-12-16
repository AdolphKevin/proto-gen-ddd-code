package proto_generate

import "testing"

func TestGenDTO(t *testing.T) {
	dataList, dataMap, err := Load("../example.proto")
	if err != nil {
		t.Error(err)
		return
	}
	err = GenDTO(dataList, dataMap, "./out.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
