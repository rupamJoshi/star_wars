package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	DBName         string `yaml:"db_name"`
	CollectionName string `yaml:"collection_name"`
}

type SWAPIConfig struct {
	SWAPIURLPeople string `yaml:"swapi_url_people"`
}

type Config struct {
	Host        string      `yaml:"host"`
	Port        string      `yaml:"port"`
	Schema      string      `yaml:"schema"`
	DBConfig    DBConfig    `yaml:"db_config"`
	SWAPIConfig SWAPIConfig `yaml:"swapi_config"`
}

func readConfig(path string) *Config {

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v", config)
	return config
}

func NewConfig(path string) *Config {
	return readConfig(path)
}
