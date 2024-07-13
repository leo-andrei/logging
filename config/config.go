package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DriverType string `json:"driver_type"`
	FilePath   string `json:"file_path"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
