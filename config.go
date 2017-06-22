package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Db struct {
		Type             string `yaml:"type"`
		ConnectionString string `yaml:"connectionstring"`
	}
}

func (c *Config) Load() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("cant load config file   %#v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %#v", err)
	}
}
