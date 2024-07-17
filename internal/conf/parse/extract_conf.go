package parse

import (
	"log"
	"os"

	"github.com/zakhaev26/canario/pkg/versioning"
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

	if config.Version != versioning.GetSDKVersion() {
		log.Fatal("SDK Version mismatch. Please try to rectify the SDK Version in your `canario.yml` file. The SDK Version is ", versioning.GetSDKVersion())
	}

	return config
}
