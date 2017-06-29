package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func main() {
	config := &Config{}
	config.Load()
	sqler := getSqler(config)

	v := sqler.Version()
	fmt.Printf("current version: %v\n", v)
	if v < 0 {
		prepareDb(sqler)
	}

	files, err := ioutil.ReadDir(config.Db.Semillas)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		i := int(strings.Index(file.Name(), "."))
		if i > v {
			semillas := ReadYaml(config.Db.Semillas + "/" + file.Name())
			sqler.Update(semillas)
		}
	}
}

func getSqler(config *Config) Sqler {
	var sqler Sqler

	switch config.Db.Type {
	case "mssql":
		sqler = &MssqlSqler{}
	case "sqlite3":
		sqler = &Sqliter{}
	}

	return sqler
}

func prepareDb(sqler Sqler) {
	fmt.Printf("preparing database\n")

	init := []Semilla{}
	inittext := fmt.Sprintf(`---

- create_table: 
    name: semillas
    columns:
    - int: version
    - string: updated_at

- insert:
    table: semillas
    columns: [version, updated_at]
    values:
    - [0, '%s']`, time.Now())

	fmt.Printf("%s\n", inittext)
	err := yaml.Unmarshal([]byte(inittext), &init)
	if err != nil {
		log.Fatal(err)
	}

	sqler.Update(init)
}

func ReadYaml(filename string) []Semilla {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   %#v ", err)
	}
	m := []Semilla{}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
	return m
}
