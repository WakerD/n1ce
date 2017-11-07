package system

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func init() {
	fmt.Println("configure init begin\n")
	loadConfig()
}

type Configs struct {
	Mode    string `yaml:"mode"`
	Debug   Config `yaml:"debug"`
	Release Config `yaml:"release"`
	Test    Config `yaml:"test"`
}

type Config struct {
	Redis    RedisConfig
	Database DababaseConfig
}

type RedisConfig struct {
	Host string
	Port uint16
}

type DababaseConfig struct {
	Host     string
	Port     uint16
	Name     string
	User     string
	Password string
}

var config *Config

func loadConfig() {
	configs := Configs{}
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
