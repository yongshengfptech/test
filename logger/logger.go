package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.SugaredLogger

func Initialize(mode string, filename string) {
	// setup writerSyncer (with log file)
	writerSyncer := GetLogWriter(filename)
	// setup encoder
	encoder := SetEncoder(mode)
	// setup core by assign encoder & writerSyncer, and set debug level
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	// declare logger with preset core and add caller
	// caller is used to mention the filename and line number of each log msg
	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}

func GetLogWriter(filename string) zapcore.WriteSyncer {
	// declare log file with configurtion
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // 10 MB
		MaxBackups: 5,  // max 5 old log file will remain
		MaxAge:     30, // max days to remain old log file
		Compress:   false,
	}

	// convert io.writer to writerSyncer
	return zapcore.AddSync(lumberJackLogger)
}

func SetEncoder(mode string) zapcore.Encoder {

	// setup configuration of encoder
	var encoderConfig zapcore.EncoderConfig

	if mode == "PRODUCTION" {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.LevelKey = "level"   // exp: info, error, fatal
		encoderConfig.FunctionKey = "func" // msg from which func
		encoderConfig.MessageKey = "msg"   // detail of msg
		encoderConfig.CallerKey = "caller" // filename and line num of msg occur
		encoderConfig.StacktraceKey = "stacktrace"
	}

	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// setup encoder with preset configuration
	var encoder zapcore.Encoder

	if mode == "PRODUCTION" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}
