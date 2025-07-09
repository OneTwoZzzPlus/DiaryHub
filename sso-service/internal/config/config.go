package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env:"STORAGE_PATH" env-default:"host=localhost port=5432 user=admin password=admin dbname=postgres sslmode=disable"`
	TokenTTL    time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-default:"60m"`
	GRPC        `yaml:"grpc"`
}

type GRPC struct {
	Port    int           `yaml:"port" env:"GRPC_PORT" env-default:"8888"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}
	cfg := new(Config)
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatalf("cannot read config file: %s because %s", configPath, err)
	}
	return cfg
}
