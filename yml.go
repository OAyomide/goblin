package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

func (c *Config) readYml() *Config {
	yamlFile, err := ioutil.ReadFile("bot.config.yml")

	if err != nil {
		log.Printf("Error opening config file: %s\n\n", err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Printf("Couldn't marshal config file:: %s\n", err)
	}
	return c
}

func getToken() string {
	var c Config

	c.readYml()

	v, err := json.Marshal(c)

	//if there was an error Marshaling our struct to json
	if err != nil {
		log.Printf("Error marshalling our json file:: %s\n", err)
	}

	return string(v)
}
