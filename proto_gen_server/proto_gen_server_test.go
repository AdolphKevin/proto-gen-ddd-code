package proto_gen_server

import (
	"testing"
)

func TestGenServer(t *testing.T) {
	err := GenServer("../test.proto", "./output.txt")
	if err != nil {
		t.Error(err)
		return
	}
}
