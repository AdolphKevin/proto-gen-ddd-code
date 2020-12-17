package proto_generate

import "testing"

func TestGenDTO(t *testing.T) {
	dataList, err := Load("../example.proto")
	if err != nil {
		t.Error(err)
		return
	}
	err = GenDTO(dataList, "../gen_result/dto.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
