package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const localConfigPath = "./config/local.yml"

type Config struct {
	Env         string        `yaml:"env"          env:"ENV"          env-default:"dev"`
	StoragePath string        `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl"    env:"TOKEN_TTL"    env-default:"60m"`
	GRPC        `yaml:"grpc"`
	REST        `yaml:"rest"`
	SMTP        `yaml:"smtp"`
}

type GRPC struct {
	Port    int           `yaml:"port"    env:"GRPC_PORT"    env-default:"7070"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"60s"`
}

type REST struct {
	Port int    `yaml:"port" env:"REST_PORT" env-default:"8080"`
	Cors string `yaml:"cors" env:"REST_CORS" env-default:"*"`
}

type SMTP struct {
	Host     string `yaml:"host"     env:"SMTP_HOST"     env-required:"true"`
	Addr     string `yaml:"addr"     env:"SMTP_ADDR"     env-required:"true"`
	Username string `yaml:"username" env:"SMTP_USERNAME" env-required:"true"`
	Password string `yaml:"password" env:"SMTP_PASSWORD" env-required:"true"`
	Sender   string `yaml:"sender"   env:"SMTP_SENDER"   env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Printf("[ERROR] CONFIG_PATH is not set! Using '%s'.\n", localConfigPath)
		configPath = localConfigPath
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}
	cfg := new(Config)
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("Cannot read config file: %s because %s", configPath, err)
	}
	return cfg
}
