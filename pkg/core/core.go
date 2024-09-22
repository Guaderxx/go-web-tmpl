package core

import (
	"context"

	"github.com/Guaderxx/gowebtmpl/config"
	"github.com/Guaderxx/gowebtmpl/ent"
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/aviper"
)

type Core struct {
	Ctx     context.Context
	Env     string // Execute environment
	Options Options
	Config  *config.Config
	Logger  alog.ALogger
	DB      *ent.Client
}

func New(cfg config.Config, ops Options) (*Core, error) {
	var err error
	c := Core{
		Ctx:     context.Background(),
		Env:     cfg.Env,
		Config:  &cfg,
		Options: ops,
	}

	c.initLog()

	c.initDB()

	return &c, err
}

func (c *Core) Close() {
	c.closeDB()
}

// WithContext  Replace context
func (c *Core) WithContext(ctx context.Context) {
	c.Ctx = ctx
}

func NewByConfig(loc string, ops Options) (*Core, error) {
	var cfg config.Config
	err := aviper.Load(loc, &cfg)
	if err != nil {
		return nil, err
	}
	c, err := New(cfg, ops)
	return c, err
}
