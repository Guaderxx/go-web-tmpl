package aviper

import "github.com/spf13/viper"

var (
	// establishing defaults
	SetDefault = viper.SetDefault

	// read config files
	SetConfigName = viper.SetConfigName
	SetConfigType = viper.SetConfigType // REQUIRED if the config file does not have the extension in the name
	AddConfigPath = viper.AddConfigPath // can call multiple times to add many search paths
	ReadInConfig  = viper.ReadInConfig

	// write config files
	WriteConfig       = viper.WriteConfig // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	SafeWriteConfig   = viper.SafeWriteConfig
	WriteConfigAs     = viper.WriteConfigAs
	SafeWriteConfigAs = viper.SafeWriteConfigAs // will error since it has already been written

	// read config from io.Reader
	ReadConfig = viper.ReadConfig

	// set override
	Set = viper.Set

	// work with environment variables
	AutomaticEnv      = viper.AutomaticEnv
	BindEnv           = viper.BindEnv
	SetEnvPrefix      = viper.SetEnvPrefix
	SetEnvKeyReplacer = viper.SetEnvKeyReplacer
	AllowEmptyEnv     = viper.AllowEmptyEnv

	// remote key/value store support
	// to enable remote support, do a blank import of the viper/remote
	// import _ "github.com/spf13/viper/remote"
	// ...

	// get values
	Get                = viper.Get
	GetBool            = viper.GetBool
	GetFloat64         = viper.GetFloat64
	GetInt             = viper.GetInt
	GetIntSlice        = viper.GetIntSlice
	GetString          = viper.GetString
	GetStringMap       = viper.GetStringMap
	GetStringMapString = viper.GetStringMapString
	GetStringSlice     = viper.GetStringSlice
	GetTime            = viper.GetTime
	GetDuration        = viper.GetDuration
	IsSet              = viper.IsSet
	AllSettings        = viper.AllSettings

	// unmarshaling
	Unmarshal    = viper.Unmarshal
	UnmarshalKey = viper.UnmarshalKey
)

// errors
type ConfigFileNotFoundError = viper.ConfigFileNotFoundError
