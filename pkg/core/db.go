package core

import (
	"fmt"

	"github.com/Guaderxx/gowebtmpl/pkg/db"
)

func (c *Core) initDB() {
	var err error

	if c.Options.DB {
		c.DB, err = db.Open(fmt.Sprintf(
			"postgresql://%s:%s@%s:%d/%s",
			c.Config.DB.User,
			c.Config.DB.Password,
			c.Config.DB.Host,
			c.Config.DB.Port,
			c.Config.DB.Schema,
		))
		if err != nil {
			c.Logger.Fatal("init db failed", "error", err.Error())
		}
		c.Logger.Info("init db succeed")
	}
}

func (c *Core) closeDB() {
	c.DB.Close()
}
