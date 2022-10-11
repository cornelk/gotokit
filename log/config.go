package log

import (
	"fmt"

	"github.com/cornelk/gotokit/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config represents configuration for a logger.
type Config = zap.Config

// ConfigForEnv returns the default config for the given environment.
// The returned config can be adjusted and used to create a logger with
// custom config using the NewWithConfig() function.
func ConfigForEnv(environment env.Environment) (Config, error) {
	var conf Config

	switch environment {
	case env.Test, env.Development:
		conf = zap.NewDevelopmentConfig()
		conf.Encoding = "console"

	case env.Staging, env.Production:
		conf = zap.NewProductionConfig()
		conf.Encoding = "json"

	default:
		return Config{}, fmt.Errorf("invalid environment specified '%v'", environment)
	}

	conf.DisableCaller = true
	conf.DisableStacktrace = true
	conf.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	conf.Sampling = nil
	return conf, nil
}
