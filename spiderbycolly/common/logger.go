package common

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"time"
)

var Logger *zap.SugaredLogger

func InitLogger(name string) {
	// 已完成初始化
	if Logger != nil {

		return
	}

	filename := fmt.Sprintf("./logs/%s.log", name)

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "linenum",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(getWriter(filename)), zapcore.AddSync(os.Stdout)),
		atomicLevel)

	caller := zap.AddCaller()
	development := zap.Development()

	logger := zap.New(core, caller, development)
	Logger = logger.Sugar()
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05]"))
}

func getWriter(filename string) io.Writer {

	name := strings.Replace(filename, ".log", "", -1) + "_%Y-%m-%d.log"
	hook, err := rotatelogs.New(name,
		rotatelogs.WithRotationCount(8),           // 文件最大保存份数
		rotatelogs.WithRotationTime(24*time.Hour), //文件切割时间
	)

	if err != nil {

		panic(err)
	}

	return hook
}
