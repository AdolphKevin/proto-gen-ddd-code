package mysql_generate

import "testing"

func TestGenerate(t *testing.T) {
	tables, err := Load("../create_table.sql")
	if err != nil {
		t.Error(err)
		return
	}

	err = GenModel(tables, "../gen_result/mysql_model")
	if err != nil {
		t.Error(err)
		return
	}

}
