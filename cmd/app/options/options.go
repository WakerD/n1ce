package options

import (
	"github.com/gin-gonic/gin"
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
	configs := loadConfig()
	s := ServerRunOptions{
		Mongodb: configs.Mongo,
		Redis:   configs.Redis,
	}
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Mongodb.Host, "mongodb-host", s.Mongodb.Host,
		"If non-empty, use secure SSH proxy to the nodes, using this user keyfile")
}

func loadConfig() config {
	configs := configs{}
	b, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil
	}
	err = yaml.Unmarshal(b, &configs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	switch gin.Mode {
	case "debug":
		config := configs.Debug
	case "release":
		config := configs.Release
	case "test":
		config := configs.Test
	default:
		log.Panic("Unkown mode")
	}
	return config
}
