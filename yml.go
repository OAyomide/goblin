package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

type MessageStruct struct {
	Welcome struct {
		typing bool
		text   string
	}
}

/**
* Here, we wan to parse the yml file retrieve our specified content file
* We will then pass this to appropriate functions. For example: Reply a user
**/

func (c *MessageStruct) parseContentFile() *MessageStruct {
	contentFile, err := ioutil.ReadFile("content.yml")
	if err != nil {
		log.Printf("Error opening content file: %s\n\n", err)
		panic(err)
	}

	er := yaml.Unmarshal(contentFile, c)

	if er != nil {
		log.Printf("Couldnt marshal content file: %s\n\n", er)
	}
	return c
}

func getContents() string {
	var c MessageStruct

	c.parseContentFile()

	v, err := json.Marshal(c)

	if err != nil {
		log.Printf("Error marshalling our json file: %s\n", err)
	}
	return string(v)
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
	//print out the parse file <@for debug only?@>
	fmt.Printf("Here is the parsed content.yml: %s\n\n", c)
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
