package baseutils

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func Test_LogInit(t *testing.T) {
	InitLogger("", 1, 5, 30, zapcore.DebugLevel, true)
	Log.Info("format")
}

func Test_LogDefault(t *testing.T) {
	df := new(defaultLogger)
	InitLog(df)
	Log.Info("format")
}
