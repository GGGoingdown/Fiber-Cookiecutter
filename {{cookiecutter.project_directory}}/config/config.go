package config

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"

	"github.com/GGGoingdown/Fiber-Cookiecutter/utils"
)

type Config struct {
	Mode          string `mapstructure:"MODE" validate:"omitempty,oneof=development production test"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" validate:"omitempty"`

	// Log
	LogLevel string `mapstructure:"LOG_LEVEL" validate:"omitempty,oneof=debug info warn error"`
	LogPath  string `mapstructure:"LOG_PATH" validate:"omitempty,dirpath"`

	// Sentry
	SentryDsn             string  `mapstructure:"SENTRY_DSN" validate:"omitempty"`
	SentryTraceSampleRate float64 `mapstructure:"SENTRY_TRACE_SAMPLE_RATE" validate:"omitempty"`
}

func (c *Config) setDefaultValue() {
	if c.Mode == "" {
		c.Mode = "development"
	}

	if c.ServerAddress == "" {
		c.ServerAddress = ":8080"
	}

	if c.LogLevel == "" {
		c.LogLevel = "debug"
	}

	if c.LogPath == "" {
		c.LogPath = "./storage/logs"
	}

	if c.SentryDsn != "" {
		log.Println("Sentry is enabled")
		if c.Mode == "production" && c.SentryTraceSampleRate == 0 {
			c.SentryTraceSampleRate = 0.1
		}
		if c.SentryTraceSampleRate == 0 {
			c.SentryTraceSampleRate = 1
		}

	}
}

func (c *Config) ToZapLogLevel() zapcore.Level {
	switch c.LogLevel {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}

func (c *Config) Show() {
	utils.PrintPretty(c)
}

func NewConfig(path string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	errors := utils.StructValidator(config)
	if errors != nil {
		return nil, errors[0]
	}

	config.setDefaultValue()

	config.Show()

	return &config, nil
}
