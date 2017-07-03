package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db struct {
		Semillas         string `yaml:"semillas"`
		Type             string `yaml:"type"`
		ConnectionString string `yaml:"connectionstring"`
	}
}

func (c *Config) Load(filePath string) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("cant load config file   %#v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
}
