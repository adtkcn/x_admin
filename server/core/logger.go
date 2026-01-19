package core

import (
	"os"
	"time"
	"x_admin/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger = NewLogger(config.LogConfig)

func NewLogger(logConfig config.LogConfigStruct) *zap.SugaredLogger {
	// 1. 根据配置决定 Encoder
	var encoder zapcore.Encoder
	if logConfig.EnableConsole {
		// 如果开启了控制台（通常意味着开发环境），使用易读的格式
		encoder = getConsoleEncoder()
	} else {
		// 否则使用 JSON 格式（生产环境推荐）
		encoder = getJSONEncoder()
	}

	// 2. 根据配置动态组合输出目标 (核心优化点)
	var writers []zapcore.WriteSyncer

	// 判断是否输出到控制台
	if logConfig.EnableConsole {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}

	// 判断是否输出到文件
	if logConfig.EnableFile {
		fileWriter := getLumberjackWriter(logConfig)
		writers = append(writers, fileWriter)
	}

	// 如果都没有开启，至少输出到控制台防止静默失败
	if len(writers) == 0 {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}

	// 使用 MultiWriteSyncer 合并多个输出
	var writeSyncer zapcore.WriteSyncer
	if len(writers) == 1 {
		writeSyncer = writers[0] // 只有一个直接用
	} else {
		writeSyncer = zapcore.NewMultiWriteSyncer(writers...) // 多个合并
	}

	// 3. 解析日志级别
	level := parseLogLevel(logConfig.Level)

	// 4. 构建 Core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 5. 创建 Logger
	logger := zap.New(core, zap.AddCaller())

	return logger.Sugar()
}

// getConsoleEncoder 控制台友好格式 (带颜色)
func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getJSONEncoder 生产环境 JSON 格式
func getJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLumberjackWriter 创建文件切割写入器
func getLumberjackWriter(logConfig config.LogConfigStruct) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.Filename,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
		Compress:   logConfig.Compress,
	})
}

// 自定义时间格式
var customTimeEncoder = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// parseLogLevel 将字符串转换为 zapcore.Level
func parseLogLevel(levelStr string) zapcore.Level {
	switch levelStr {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
