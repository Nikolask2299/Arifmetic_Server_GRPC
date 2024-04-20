package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
)

type Config struct {
	Server string  `json:"server" yaml:"server"`
	HttpPort string `json:"httpPort" yaml:"httpPort"`
	PortData string 	 	`json:"portData" yaml:"portData"`
	PortAuth string `json:"portAuth" yaml:"portAuth"`
	Timeout string `json:"timeout" yaml:"timeout"`
	TockenTTL string `json:"tockenTTL" yaml:"tockenTTL"` 
	StoragePath string  `json:"storage_path" yaml:"storage_path"`
	CountAgent int  	 	`json:"count_agent" yaml:"count_agent"`
	PathMain string `json:"html_path_main" yaml:"html_path_main"`
	PathAuth string `json:"html_path_auth" yaml:"html_path_auth`
}

 func Mustload() *Config {
	path := fethConfigPath()
	if path == "" {
		panic("Failed to load config file")
	}
	return MustloadPath(path)
 }

 func MustloadPath(path string) *Config {
	var config Config

	file, err := os.Open(path) 
	if err != nil {
		panic("Failed to open config file")
	}
	defer file.Close()
	bufer := make([]byte, 0, 255)
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		bytes := reader.Bytes()
		bufer = append(bufer, bytes...)
	}

	err = json.Unmarshal(bufer, &config)
	if err != nil {
		panic("Failed to unmarshal config file" + err.Error())
	}

	return &config
 }

 func fethConfigPath() string {
	var result string

	flag.StringVar(&result, "config", "", "Path to config file")
	flag.Parse()

	if result == "" {
		panic("Failed to load config file")
	}

	return result
 }