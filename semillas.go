package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	var sqler Sqler
	config := &Config{}
	config.Load()
	files, err := ioutil.ReadDir(config.Db.Semillas)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		semillas := ReadYaml(config.Db.Semillas + "/" + file.Name())

		switch config.Db.Type {
		case "mssql":
			sqler = &MssqlSqler{}
		case "sqlite3":
			sqler = &Sqliter{}
		}

		sqler.Update(semillas)
	}
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
