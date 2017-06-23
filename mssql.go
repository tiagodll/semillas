package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb"
)

type MssqlSqler struct{}

func (this *MssqlSqler) Migrate(migrations []Migration) {
	config := &Config{}
	config.Load()
	db, err := sql.Open(config.Db.Type, config.Db.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, m := range migrations {
		fmt.Printf("%s", this.ToSql(&m))
		result, err := db.Exec(this.ToSql(&m))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", result)
	}
}

func (this *MssqlSqler) ToSql(m *Migration) string {
	switch {
	case m.IsCreateTable.Name != CreateTable{}.Name:
		return this.createTableToSql(&m.IsCreateTable)
	case m.IsDropTable != DropTable{}:
		return this.dropTableToSql(&m.IsDropTable)
	case m.IsAddColumn != AddColumn{}:
		return this.addColumnToSql(&m.IsAddColumn)
	case m.IsRemoveColumn != RemoveColumn{}:
		return this.removeColumnToSql(&m.IsRemoveColumn)
	case m.IsInsert.Table != Insert{}.Table:
		return this.insertToSql(&m.IsInsert)
	}
	return m.SQL
}

func (this *MssqlSqler) createTableToSql(t *CreateTable) string {

	cmd := fmt.Sprintf(`CREATE TABLE %s (`, t.Name)
	for _, c := range t.Columns {
		cname, ctype := this.GetNameAndType(&c)
		cmd = fmt.Sprintf("%s%s %s, ", cmd, cname, ctype)
	}
	cmd = fmt.Sprintf("%s);\n", cmd[0:len(cmd)-2])
	return cmd
}

func (this *MssqlSqler) GetNameAndType(c *Column) (string, string) {
	switch {
	case c.StringColumn != "":
		return c.StringColumn, "[varchar](max)"
	case c.BoolColumn != "":
		return c.BoolColumn, "[bit]"
	case c.IntColumn != "":
		return c.IntColumn, "[int]"
	case c.FloatColumn != "":
		return c.FloatColumn, "[float]"
	}
	return "", ""
}

func (this *MssqlSqler) dropTableToSql(t *DropTable) string {
	cmd := fmt.Sprintf("DROP TABLE %s;\n", t.Name)
	return cmd
}

func (this *MssqlSqler) addColumnToSql(c *AddColumn) string {
	return ""
}

func (this *MssqlSqler) removeColumnToSql(c *RemoveColumn) string {
	return ""
}

func (this *MssqlSqler) insertToSql(c *Insert) string {
	resp := ""
	column_names := ""
	for _, col := range c.Columns {
		column_names = fmt.Sprintf("%s, %s", column_names, col)
	}
	for _, row := range c.Values {
		values := ""
		for _, cell := range row {
			switch reflect.ValueOf(cell).Kind() {
			case reflect.String:
				values = fmt.Sprintf("%s, '%s'", values, cell)
			case reflect.Invalid:
				values = fmt.Sprintf("%s, NULL", values)
			default:
				values = fmt.Sprintf("%s, %v", values, cell)
			}
		}
		resp += fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);\n", c.Table, column_names[2:len(column_names)], values[2:len(values)])
	}
	return resp
}
