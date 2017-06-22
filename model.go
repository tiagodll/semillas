package main

type Sqler interface {
	ToSql(m *Migration) string
}

type Migration struct {
	IsCreateTable  CreateTable  `yaml:"create_table"`
	IsDropTable    DropTable    `yaml:"drop_table"`
	IsAddColumn    AddColumn    `yaml:"add_column"`
	IsRemoveColumn RemoveColumn `yaml:"remove_column"`
	SQL            string       `yaml:"sql"`
}

type CreateTable struct {
	Sqler
	Name    string   `yaml:"name"`
	Columns []Column `yaml:"columns"`
}

type Column struct {
	Sqler
	StringColumn string `yaml:"string"`
	BoolColumn   string `yaml:"bool"`
	IntColumn    string `yaml:"int"`
	FloatColumn  string `yaml:"float"`
	DateColumn   string `yaml:"date"`
}

type DropTable struct {
	Sqler
	Name string `yaml:"name"`
}

type AddColumn struct {
	Sqler
	Name string `yaml:"name"`
}

type RemoveColumn struct {
	Sqler
	Name string `yaml:"name"`
}