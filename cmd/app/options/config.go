package options

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type configs struct {
	Debug   config `yaml:"debug"`
	Release config `yaml:"release"`
	Test    config `yaml:"test"`
}

type config struct {
	Mongo mongodbConfig
	Redis redisConfig
}

type mongodbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type redisConfig struct {
	Port string
}

var config *Config

func loadConfig() {
	configs := configs{}
	b, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	switch configs.Mode {
	case "debug":
		config = &configs.Debug
	case "release":
		config = &configs.Release
	case "test":
		config = &configs.Test
	default:
		panic(fmt.Sprintf("Unkown mode %s", configs.Mode))
	}
}

func GetConfig() *Config {
	return config
}
