package logger

import (
	"go-ecommerce-backend-api/pkg/setting"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel
	// debug level
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName, //"//var/log/myapp/foo.log",
		MaxSize:    config.MaxSize,     // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1716714967.877995 -> 2024-05-26T16:16:07.877+0700
	// ts -> Time
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "time"

	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.go:24"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile(
		"./log/log.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	if err != nil {
		panic(err)
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
