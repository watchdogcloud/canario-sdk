package conf

import "github.com/zakhaev26/canario/internal/conf/parse"

func CreateNewConf() parse.Config {
	cfg := parse.ExtractYAML()
	cfg.SetDefaultsIfFieldsMissing()
	return cfg
}
