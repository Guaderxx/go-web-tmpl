package config

import (
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/web"
)

type Config struct {
	Env     string
	AppName string `mapstructure:"app_name"`
	Log     alog.Options
	Web     web.Options
	DB      DB
}

type DB struct {
	Name     string
	Host     string
	Port     int
	User     string
	Schema   string
	Password string
}
