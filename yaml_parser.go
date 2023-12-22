package yaml_parser

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host    string
	Port    int
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

func ParseYamlFile(filePath string) (Config, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	return ParseYaml(b)
}

func ParseYaml(b []byte) (Config, error) {
	b = []byte(os.ExpandEnv(string(b)))

	var cfg Config
	err := yaml.Unmarshal(b, &cfg)

	return cfg, err
}
