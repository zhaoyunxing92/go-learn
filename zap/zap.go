package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/natefinch/lumberjack"
)

func main() {

	encoder := zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "time",
			CallerKey:     "line",
			NameKey:       "logger",
			FunctionKey:   "func",
			StacktraceKey: "stacktrace",
			//EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.0000"),
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		})
	//colorable.NewColorableStdout()
	// 设置日志级别
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   ".logs/zap.log", //日志文件存放目录
		MaxSize:    1,               //文件大小限制,单位MB
		MaxBackups: 5,               //最大保留日志文件数量
		MaxAge:     30,              //日志文件保留天数
		Compress:   true,            //是否压缩处理
	})

	//jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
	//	MessageKey:     "msg",
	//	LevelKey:       "level",
	//	TimeKey:        "time",
	//	CallerKey:      "line",
	//	NameKey:        "logger",
	//	FunctionKey:    "func",
	//	StacktraceKey:  "stacktrace",
	//	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
	//	EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
	//	EncodeCaller:   zapcore.ShortCallerEncoder,
	//	EncodeDuration: zapcore.SecondsDurationEncoder,
	//})

	fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(file, zapcore.AddSync(os.Stdout)), level)
	log := zap.New(fileCore, zap.AddCaller()).Sugar()

	for {
		log.Info("log level is info")
		log.Error("log level is error")
		log.Warn("log level is warn")
		log.Debug("log level debug")
	}
}
