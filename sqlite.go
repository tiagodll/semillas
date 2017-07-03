package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"semillas/config"

	_ "github.com/mattn/go-sqlite3"
)

type Sqliter struct {
	config config.Config
}

func (this *Sqliter) Version() int {
	db, err := sql.Open(this.config.Db.Type, this.config.Db.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Printf("config: %v\n", db)

	row, err := db.Query("Select version from semillas")
	if err != nil {
		return -1
	}
	version := 0
	row.Scan(&version)

	return version
}

func (this *Sqliter) Init() {
	os.Remove(this.config.Db.ConnectionString)
}

func (this *Sqliter) Update(semillas []Semilla) {

	db, err := sql.Open(this.config.Db.Type, this.config.Db.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, s := range semillas {
		fmt.Printf("%s", this.ToSql(&s))
		result, err := db.Exec(this.ToSql(&s))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", result)
	}
}

func (this *Sqliter) ToSql(s *Semilla) string {
	switch {
	case s.IsCreateTable.Name != CreateTable{}.Name:
		return this.createTableToSql(&s.IsCreateTable)
	case s.IsDropTable != DropTable{}:
		return this.dropTableToSql(&s.IsDropTable)
	case s.IsAddColumn != AddColumn{}:
		return this.addColumnToSql(&s.IsAddColumn)
	case s.IsRemoveColumn != RemoveColumn{}:
		return this.removeColumnToSql(&s.IsRemoveColumn)
	case s.IsInsert.Table != Insert{}.Table:
		return this.insertToSql(&s.IsInsert)
	}
	return s.SQL
}

func (this *Sqliter) createTableToSql(t *CreateTable) string {

	cmd := fmt.Sprintf(`CREATE TABLE %s (`, t.Name)
	for _, c := range t.Columns {
		cname, ctype := this.GetNameAndType(&c)
		cmd = fmt.Sprintf("%s%s %s, ", cmd, cname, ctype)
	}
	cmd = fmt.Sprintf("%s);\n", cmd[0:len(cmd)-2])
	return cmd
}

func (this *Sqliter) GetNameAndType(c *Column) (string, string) {
	switch {
	case c.StringColumn != "":
		return c.StringColumn, "text"
	case c.BoolColumn != "":
		return c.BoolColumn, "int"
	case c.IntColumn != "":
		return c.IntColumn, "int"
	case c.FloatColumn != "":
		return c.FloatColumn, "float"
	}
	return "", ""
}

func (this *Sqliter) dropTableToSql(t *DropTable) string {
	cmd := fmt.Sprintf("DROP TABLE %s;\n", t.Name)
	return cmd
}

func (this *Sqliter) addColumnToSql(c *AddColumn) string {
	return ""
}

func (this *Sqliter) removeColumnToSql(c *RemoveColumn) string {
	return ""
}

func (this *Sqliter) insertToSql(c *Insert) string {
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
