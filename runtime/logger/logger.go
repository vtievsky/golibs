package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateZapLogger(debug, stacktrace bool) *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	encoder := zapcore.NewJSONEncoder(cfg)

	stacktraceLevel := zapcore.ErrorLevel
	if !stacktrace {
		stacktraceLevel = zapcore.FatalLevel + 1
	}

	stdoutFiler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		if debug {
			return level < zapcore.ErrorLevel
		}

		return level > zapcore.DebugLevel && level < zapcore.ErrorLevel
	})

	stderrFilter := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stdout),
			stdoutFiler,
		),
		zapcore.NewCore(
			encoder,
			zapcore.Lock(os.Stderr),
			stderrFilter,
		),
	)

	return zap.New(core, zap.AddStacktrace(stacktraceLevel))
}
