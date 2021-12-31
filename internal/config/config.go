package config

import cfg "github.com/dollarkillerx/common/pkg/config"

// config ...
type config struct {
	LoggerConfig       cfg.LoggerConfig
	EnablePlayground   bool
	JWTConfig          JWTConfiguration
	Debug              bool
	CORSAllowedOrigins []string
}

type JWTConfiguration struct {
	SecretKey    string
	OperationKey string
}

// Global ...
var conf config

func InitConfig(configName string, configPaths []string) error {
	return cfg.InitConfiguration(configName, configPaths, &conf)
}

func GetLoggerConfig() cfg.LoggerConfig {
	return conf.LoggerConfig
}

func GetEnablePlayground() bool {
	return conf.EnablePlayground
}

func GetJWTConfig() JWTConfiguration {
	return conf.JWTConfig
}

func GetCORSAllowedOrigins() []string {
	return conf.CORSAllowedOrigins
}
