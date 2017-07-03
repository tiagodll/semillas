package main

import (
	"semillas/config"

	_ "github.com/denisenkom/go-mssqldb"
)

type MssqlSqler struct {
	config config.Config
}

func (this *MssqlSqler) Version() int { return -1 }

func (this *MssqlSqler) Init() {}

func (this *MssqlSqler) Update(migrations []Semilla) {}

func (this *MssqlSqler) ToSql(m *Semilla) string { return "" }

func (this *MssqlSqler) createTableToSql(t *CreateTable) string { return "" }

func (this *MssqlSqler) GetNameAndType(c *Column) (string, string) { return "", "" }

func (this *MssqlSqler) dropTableToSql(t *DropTable) string { return "" }

func (this *MssqlSqler) addColumnToSql(c *AddColumn) string { return "" }

func (this *MssqlSqler) removeColumnToSql(c *RemoveColumn) string { return "" }

func (this *MssqlSqler) insertToSql(c *Insert) string { return "" }
