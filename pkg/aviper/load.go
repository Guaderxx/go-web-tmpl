package aviper

import (
	"path/filepath"

	"github.com/Guaderxx/gowebtmpl/config"
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
)

func Load(loc string, config *config.Config) error {
	dir, file := filepath.Split(loc)
	SetConfigName(file)
	fileType := filepath.Ext(file)
	SetConfigType(fileType[1:])

	AddConfigPath(dir)
	// optionally look for config in the working directory
	AddConfigPath(".")

	err := ReadInConfig()
	if err != nil {
		if _, ok := err.(ConfigFileNotFoundError); ok {
			alog.Fatal("load config", "location", loc, "error", "not found")
		} else {
			alog.Fatal("load config", "location", loc, "error", err)
		}
	}

	err = Unmarshal(config)
	return err
}
