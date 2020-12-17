package mysql_generate

import "testing"

func TestGenerate(t *testing.T) {
	tables, err := Load("../create_table.sql")
	if err != nil {
		t.Error(err)
		return
	}

	err = Generate(tables)
	if err != nil {
		t.Error(err)
		return
	}

}
