package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

type Config struct {
	Port uint64 `json:"port" validate:"required"`
}

func NewConfig() (Config, error) {
	const path = "service_test.json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("Invalid read config file[%s]: %w", path, err)
	}

	var cfg Config
	if err := jsoniter.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("Invalid unmarshal config file[%s]: %w", path, err)
	}

	if err := validator.New().Struct(cfg); err != nil {
		return Config{}, fmt.Errorf("Invalid config[%s]: %w", path, err)
	}

	return cfg, nil
}
