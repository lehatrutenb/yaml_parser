package yamlparser

import (
	"os"

	"gopkg.in/yaml.v2"
)

func ParseYamlFile(obj *any, filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return ParseYaml(obj, b)
}

func ParseYaml(obj *any, b []byte) error {
	b = []byte(os.ExpandEnv(string(b)))

	return yaml.Unmarshal(b, obj)
}
