package main

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {

	// 配置 sugaredLogger
	var sugaredLogger *zap.SugaredLogger
	writer := zapcore.AddSync(os.Stdout)

	// 格式相关的配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder  // 修改时间戳的格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 日志级别使用大写显示
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)  // 将日志级别设置为 DEBUG
	logger := zap.New(core, zap.AddCaller(), zap.Fields(zapcore.Field{  // 添加 uuid 字段
		Key:       "uuid",
		Type:      zapcore.StringType,
		String:    uuid.New().String(),
	}))
	sugaredLogger = logger.Sugar()

	// 打印日志
	sugaredLogger.Debugf("i am debug, using %s", "sugar")   // 这行现在可以打印出来了！
	sugaredLogger.Infof("i am info, using %s", "sugar")      // INFO  级别日志，这个会正常打印
	sugaredLogger.Warnf("i am warn, using %s", "sugar")    // WARN  级别日志，这个会正常打印
	sugaredLogger.Errorf("i am error, using %s", "sugar")   // ERROR 级别日志，这个会打印，并附带堆栈信息
	sugaredLogger.Fatalf("i am fatal, using %s", "sugar")    // FATAL 级别日志，这个会打印，附带堆栈信息，并调用 os.Exit 退出
	fmt.Println("can i be printed?")  // 这行不会打印，呃...上面已经退出了
}