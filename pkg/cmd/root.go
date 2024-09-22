package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/Guaderxx/gowebtmpl/config"
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/aviper"
	"github.com/Guaderxx/gowebtmpl/pkg/core"
	"github.com/Guaderxx/gowebtmpl/pkg/web"
	"github.com/Guaderxx/gowebtmpl/pkg/web/validate"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	cfg     config.Config

	rootCmd = &cobra.Command{
		Use:   "tmpl",
		Short: "A Web Template",
		Run:   Exec,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		alog.Fatal("execute root cmd error", "error", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config/tmpl.toml)")
	viper.SetDefault("author", "Guaderxx <guaderxx@gmail.com>")
	viper.SetDefault("license", "GPL")
}

func initConfig() {
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		cfgFile = filepath.Join(home, ".config", "tmpl.toml")
	}

	err := aviper.Load(cfgFile, &cfg)

	if err != nil {
		alog.Fatal("load config error", "error", err)
	}
	alog.Info("init config succeed", "config", cfg)
}

func Exec(cmd *cobra.Command, args []string) {
	// core
	c, err := core.New(cfg, core.Options{
		Logger: true,
		DB:     true,
	})
	if err != nil {
		alog.Panic("init core failed", "error", err)
	}
	defer c.Close()

	// binding.Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("password", validate.Password)
		if err != nil {
			c.Logger.Fatal("register validation password failed", "error", err.Error())
		}
	} else {
		c.Logger.Fatal("get binding.validator failed")
	}

	eng := web.New(c.Logger)

	web.Route(eng, c)

	srv := &http.Server{
		Addr:           ":" + c.Config.Web.Port,
		Handler:        eng,
		ReadTimeout:    time.Duration(c.Config.Web.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.Config.Web.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(c.Config.Web.IdleTimeout) * time.Second,
		MaxHeaderBytes: c.Config.Web.MaxHeaderMB << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			alog.Fatal("listen server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	alog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(c.Ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		alog.Fatal("Server Shutdown failed", "error", err)
	}

	select {
	case <-ctx.Done():
		alog.Info("timeout of 5 seconds.")
	}

	alog.Info("Server Exit Succeed")
}
