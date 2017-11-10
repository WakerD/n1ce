package options

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type ServerRunOptions struct {
	Mongodb mongodbConfig
	Redis   redisConfig
}

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

func NewServerRunOptions() *ServerRunOptions {
	configs, _ := loadConfig()
	s := &ServerRunOptions{
		Mongodb: configs.Mongo,
		Redis:   configs.Redis,
	}
	return s
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mongodb.Host, "mongodb-host", s.Mongodb.Host,
		"If non-empty, use secure SSH proxy to the nodes, using this user keyfile")
}

func loadConfig() (*config, error) {
	configs := configs{}
	b, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, errors.New("no config file")
	}
	err = yaml.Unmarshal(b, &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var cfg config
	switch gin.Mode() {
	case "debug":
		cfg = configs.Debug
	case "release":
		cfg = configs.Release
	case "test":
		cfg = configs.Test
	default:
		log.Panic("Unkown mode")
	}
	return &cfg, nil
}
