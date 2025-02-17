package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"DEBUG"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Addr        string        `yaml:"addr" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"10s"`
}

func MustLoad() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("CONFIG_PATH must be set")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read from config" + err.Error())
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read from config" + err.Error())
	}

	var config *Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Unable to parse config" + err.Error())
	}

	return config
}
