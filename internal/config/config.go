package config

import (
	"log"
	"os"
  "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
  Env string `yaml:"env" env-default:"local" env-required:"true"`
  StoragePath string `yaml:"storage_path" env-required:"true"`
  HttpServer HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
  Address string `yaml:"address" address-default:"localhost:8080"`
}

func MustLoad() *Config {
  configPath := os.Getenv("CONFIG_PATH")
  if configPath == "" {
    log.Fatal("Config path is not set :(")
  }

  if _, err := os.Stat(configPath); os.IsNotExist(err) {
    log.Fatalf("Config is not stated: %s", configPath)
  }
  
  var cfg Config

  if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
    log.Fatalf("Cannot read config: %s", configPath)
  }

  return &cfg
}
