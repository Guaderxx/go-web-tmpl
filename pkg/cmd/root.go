package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/Guaderxx/gowebtmpl/config"
	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/Guaderxx/gowebtmpl/pkg/aviper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	cfg     config.Config
	logger  alog.ALogger

	rootCmd = &cobra.Command{
		Use:   "tmpl",
		Short: "A Web Template",
		Run:   Exec,
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("execute root cmd error", "error", err)
	}
}

func init() {
	ctx := context.Background()
	initLogger(ctx)

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
		logger.Fatal("load config error", "error", err)
	}
	logger.Info("init config succeed", "config", cfg)
}

func initLogger(ctx context.Context) {
	logger = alog.New(ctx, cfg.Log.Formatter, cfg.Log.Level, cfg.Log.AddSource)
	logger.Info("init logger succeed")
}

func Exec(cmd *cobra.Command, args []string) {
	logger.Info("cobra - exec")
}
