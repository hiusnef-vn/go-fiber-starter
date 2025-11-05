package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/hiusnef-vn/go-fiber-starter/utils/configloader"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "DEBUG":
		return zapcore.DebugLevel
	case "INFO":
		return zapcore.InfoLevel
	case "WARN":
		return zapcore.WarnLevel
	case "ERROR":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func InitLogger(logLevel string, logger *lumberjack.Logger) *zap.Logger {
	level := getLogLevel(logLevel)
	encoder := getEncoderLog()

	var writers []zapcore.WriteSyncer
	writers = append(writers, zapcore.AddSync(os.Stdout))

	if logger != nil {
		writers = append(writers, zapcore.AddSync(logger))
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(writers...),
		level,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel), zap.Fields(zap.Int("pid", os.Getpid())))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func New(config *Config) *zap.Logger {
	var logg *lumberjack.Logger = nil
	if config.Enable {
		logg = &lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", config.DirPath, config.FileName),
			MaxSize:    config.MaxSize,
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackup,
			LocalTime:  config.LocalTime,
			Compress:   config.Compress,
		}
	}
	return InitLogger(config.Level, logg)
}

type loggerWrapper struct {
	cfg    *Config
	logger *zap.Logger
}

var (
	wrapper *loggerWrapper
	once    sync.Once
)

const (
	DefaultLoggingConfigPath = "etc/config/logging.yml"
	LoggingConfigPathEnv     = "HIUSNEF_LOGGING_CONFIG"
	LoggingEnvPrefixConfig   = "HIUSNEF_LOGGING"
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		wrapper = new(loggerWrapper)
		configPath := os.Getenv(LoggingConfigPathEnv)
		if configPath == "" {
			configPath = DefaultLoggingConfigPath
		}
		wrapper.cfg = configloader.LoadConfig[Config](configPath, LoggingEnvPrefixConfig)
		wrapper.logger = New(wrapper.cfg)
		wrapper.logger.Info("Logger initiated")
	})

	return wrapper.logger
}
