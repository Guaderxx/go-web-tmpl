package config

import "github.com/Guaderxx/gowebtmpl/pkg/alog"

type Config struct {
	Env     string
	AppName string `mapstructure:"app_name"`
	Log     alog.Options
}
