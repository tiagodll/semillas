package main

import (
	"fmt"
)

type Sqliter struct{}

func (this *Sqliter) ToSql(m *Migration) string {
	switch {
	case m.IsCreateTable.Name != CreateTable{}.Name:
		return this.createTableToSql(&m.IsCreateTable)
	case m.IsDropTable != DropTable{}:
		return this.dropTableToSql(&m.IsDropTable)
	case m.IsAddColumn != AddColumn{}:
		return this.addColumnToSql(&m.IsAddColumn)
	case m.IsRemoveColumn != RemoveColumn{}:
		return this.removeColumnToSql(&m.IsRemoveColumn)
	}
	return m.SQL
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
