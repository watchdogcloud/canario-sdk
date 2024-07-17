package conf

import "github.com/watchdogcloud/canario/internal/conf/parse"

func CreateNewConf() parse.Config {
	cfg := parse.ExtractYAML()
	cfg.SetDefaultsIfFieldsMissing()
	return cfg
}
