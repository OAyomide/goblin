package parser

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configgg struct {
	VerifyToken string `yaml:"verify_token"`
	AccessToken string `yaml:"access_token"`
	AppSecret   string `yaml:"app_secret"`
}

//ParseContentFile parses the file and return the data
func ParseContentFile() string {
	configFile := filepath.FromSlash("/content.yml")
	contentFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Printf("Error opening content file: %s\n", err.Error())
		panic(err)
	}

	fl, err := yaml.Marshal(contentFile)

	if err != nil {
		log.Printf("Couldn't marshal content file: %s\n", err)
	}
	return string(fl)
}

//ReadYml is the receiver function for our config struct that allows us
//read the config file and Unmarshal the data gotten from the
//config file to the struct
func (x *Configgg) ReadYml() *Configgg {
	configFile, err := filepath.Abs("./bot.config.yml")
	if err != nil {
		log.Printf("ERROR RETURNING THE ABSOLUTE PATH FOR THE CONFIG FILE")
	}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("ERROR READING THE CONFIG FILE: %s", err.Error())
		panic(err)
	}

	fl := yaml.Unmarshal(yamlFile, x)
	if fl != nil {
		log.Printf("ERROR MARSHALLING YAML FILE::%s", fl.Error())
	}

	return x
}

//GetAccessToken returns the accesstoken grabbed from the config file
func GetAccessToken() string {
	var c Configgg

	configObj := c.ReadYml()
	return configObj.AccessToken
}
