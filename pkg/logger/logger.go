package logger

import (
	"fmt"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creating a new logger and set as global
func NewLogger(config *config.Logger) {
	logLevel, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		panic(fmt.Sprintf("Unkown log level: %v", logLevel))
	}

	var cfg zap.Config
	if config.Development {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	zap.ReplaceGlobals(logger)
}

// Close deleting logger that globally added.
func Close() {
	defer zap.L().Sync()
}
