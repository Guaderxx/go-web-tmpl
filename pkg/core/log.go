package core

import "github.com/Guaderxx/gowebtmpl/pkg/alog"

func (c *Core) initLog() {
	if c.Options.Logger {
		c.Logger = alog.New(
			c.Ctx,
			c.Config.Log.Formatter,
			c.Config.Log.Level,
			c.Config.Log.AddSource,
		)
	}
}
