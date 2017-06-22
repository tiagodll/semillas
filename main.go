package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	dirname := "migrations"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		migrations := ReadYaml(dirname + "/" + file.Name())
		sqler := &MssqlSqler{}
		Migrate(migrations, sqler)
	}
}

func ReadYaml(filename string) []Migration {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   %#v ", err)
	}
	m := []Migration{}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
	return m
}
