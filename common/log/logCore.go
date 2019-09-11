package log

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logs struct {
	Default *zap.Logger
}

var Log *Logs

func Init() {
	Log = &Logs{}
	Log.Default = NewLogger(
		"golang",
		zapcore.InfoLevel,
		viper.GetInt("log_max_size"),
		viper.GetInt("log_max_backup"),
		viper.GetInt("log_max_age"),
		viper.GetBool("log_compress"),
		"local",
	)
}

func NewLogger(
	filename string,
	level zapcore.Level,
	maxSize,
	maxBackup,
	maxAge int,
	compress bool,
	serviceName string) *zap.Logger {
    core := newCore(
    	filename,
    	level,
    	maxSize,
    	maxBackup,
    	maxAge,
    	compress,
    )

	return zap.New(
		core, zap.AddCaller(), zap.Development(),
		zap.Fields(zap.String("serviceName", serviceName)),
	)
}

func newCore(
	filename string,
	level zapcore.Level,
	maxSize,
	maxBackup,
	maxAge int,
	compress bool) zapcore.Core {

	file := filename + "_" + time.Now().UTC().Format("2006-01-02") + ".log"
	hook := lumberjack.Logger{
		Filename: "./logs/" + file,
		MaxSize: maxSize,
		MaxBackups: maxBackup,
		MaxAge: maxAge,
		Compress: compress,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "linenum",
		MessageKey: "message",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)
}