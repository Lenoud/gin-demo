// utils/logger/logger.go
package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(viper.GetString("log.dir"), viper.GetString("log.filename")), // 日志路径
		MaxSize:    viper.GetInt("log.max.size"),                                               // 单个日志文件最大MB
		MaxBackups: viper.GetInt("log.max.backups"),                                            // 保留旧文件数
		MaxAge:     viper.GetInt("log.max.age"),                                                // 保留天数
		Compress:   viper.GetBool("log.compress"),
	})

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})

	// 环境切换
	env := "dev"
	var core zapcore.Core
	if env == "dev" {
		logLevel := zapcore.DebugLevel // 开发环境开启dubug
		consoleSyncer := zapcore.Lock(os.Stdout)
		// core := zapcore.NewCore(encoder, writeSyncer, logLevel)
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, logLevel),
			zapcore.NewCore(encoder, consoleSyncer, logLevel), // 控制台输出
		)
	} else {
		logLevel := zapcore.InfoLevel
		core = zapcore.NewCore(encoder, writeSyncer, logLevel)
	}

	Logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(Logger) // 让 zap.L() 使用默认 logger
	zap.L().Info("Zap logger 初始化完成",
		zap.String("logDir", viper.GetString("log.dir")),
		zap.String("logFilename", viper.GetString("log.filename")),
		zap.Int("maxSize", viper.GetInt("log.max.size")),
		zap.Int("maxBackups", viper.GetInt("log.max.backups")),
		zap.Int("maxAge", viper.GetInt("log.max.age")),
		zap.Bool("compress", viper.GetBool("log.compress")),
		zap.String("environment", env),
	)
	fmt.Println("Zap logger 初始化")
}
