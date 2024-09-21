package config

import (
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
)

type Config struct {
	Env     string
	AppName string `mapstructure:"app_name"`
	Log     alog.Options
	Web     Web
	DB      DB
	JWT     JWT
}

type DB struct {
	Name     string
	Host     string
	Port     int
	User     string
	Schema   string
	Password string
}

type Web struct {
	Port         string
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
	MaxHeaderMB  int `mapstructure:"max_header_mb"`
}

type JWT struct {
	AccessTokenSecret      string `mapstructure:"access_token_secret"`
	AccessTokenExpiryHour  int    `mapstructure:"access_token_expiry_hour"`
	RefreshTokenSecret     string `mapstructure:"refresh_token_secret"`
	RefreshTokenExpiryHour int    `mapstructure:"refresh_token_expiry_hour"`
}
