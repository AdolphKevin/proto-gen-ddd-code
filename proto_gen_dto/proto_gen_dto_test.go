package proto_gen_dto

import (
	"testing"
)

func TestGenDTO(t *testing.T) {
	err := GenDTO("../test.proto", "./output.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
