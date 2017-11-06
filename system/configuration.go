package system

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	loadConfig()
}

type Configs struct {
	Mode    string
	Debug   Config
	Release Config
	Test    Config
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
	b, err := ioutil.ReaFile("config/config.yaml")
	if err != nil {
		return err
	}
	configs := &Configs{}
	err := yaml.Unmarshal(b, &configs)
	if err != nil {
		log.fatalf("error: %v", err)
	}
	switch configs.Mode() {
	case "debug":
		config = &config.Debug
	case "release":
		config = &config.Release
	case "test":
		config = &config.Test
	default:
		panic(fmt.Sprintf("Unkown mode %s"), configs.Mode)
	}
}

func GetConfig() *Config {
	return config
}
