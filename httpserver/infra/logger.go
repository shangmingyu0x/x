package infra

import (
	"io"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"

	rotator "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
)

var (
	// BaseLogger 全局的logger
	BaseLogger *zap.Logger
	once       *sync.Once
)

// init .
func init() {
	once = new(sync.Once)
	once.Do(func() {
		encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		})

		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.WarnLevel
		})
		warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		})

		infoWriter := RotateWriter("./log/service.log")
		warnWriter := RotateWriter("./log/service.log.wf")

		core := zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
		)

		BaseLogger = zap.New(core, zap.AddCaller())
	})
}

// RotateWriter 进行日志切分
func RotateWriter(filename string) io.Writer {
	hook, err := rotator.New(
		filename+".%Y%m%d%H",
		rotator.WithLinkName(filename),
		rotator.WithMaxAge(time.Hour*24*time.Duration(7)),
		rotator.WithRotationTime(time.Hour*time.Duration(1)),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
