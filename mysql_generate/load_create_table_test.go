package mysql_generate

import "testing"

func TestLoad(t *testing.T) {
	sqlTables, err := Load("../create_table.sql")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("length : %d", len(sqlTables))
	for _, table := range sqlTables {
		t.Log(table.Name)
		for _, field := range table.Fields {
			t.Logf("name:%s,\ttype:%s,\tcommont:%s\t\n", field.Name, field.Type, field.Comment)
		}
	}
}
