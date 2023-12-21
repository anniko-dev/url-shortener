package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"600s" env-required:"true"`
}

func MustLoadConfig() *Config {
	configPath := "./config/local.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config не был найден в: %s", configPath)
	}

	var conf Config

	if err := cleanenv.ReadConfig(configPath, &conf); err != nil {
		log.Fatalf("ошибка загрузки конфигурации: %s", err)
	}

	return &conf
}
