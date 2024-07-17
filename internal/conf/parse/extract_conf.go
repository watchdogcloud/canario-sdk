package parse

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func ExtractYAML() Config {
	yml, err := os.ReadFile("./canario.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var config Config

	if err = yaml.Unmarshal([]byte(yml), &config); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(config)
	return config
}
