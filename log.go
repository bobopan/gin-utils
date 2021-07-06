package baseutils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//LogInterface log日志接口，实现get path和get log level方法
type LogInterface interface {
	GetLogPath() string
	GetLogMaxSize() int
	GetLogBackups() int
	GetLogMaxAge() int
	GetLogLevel() string
	GetLogInConsole() bool
}

type defaultLogger struct {
}

func (l defaultLogger) GetLogPath() string {
	return ""
}

func (l defaultLogger) GetLogMaxSize() int {
	return 1
}

func (l defaultLogger) GetLogBackups() int {
	return 5
}

func (l defaultLogger) GetLogMaxAge() int {
	return 30
}

func (l defaultLogger) GetLogLevel() string {
	return "DEBUG"
}

func (l defaultLogger) GetLogInConsole() bool {
	return true
}

var Log *zap.SugaredLogger
var logInterface LogInterface

func InitLogger(path string, maxSize int, maxBackups int, maxAge int, level zapcore.Level, logInConsole bool) {
	writeSyncer := getLogWriter(path, maxSize, maxBackups, maxAge, logInConsole)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	Log = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string, maxSize int, maxBackups int, maxAge int, logInConsole bool) zapcore.WriteSyncer {
	if path == "" {
		return zapcore.AddSync(os.Stdout)
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	//如果输出控制台
	if logInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

func reloadLog() error {
	logPath := logInterface.GetLogPath()
	logLevel := logInterface.GetLogLevel()
	var level zapcore.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		level = zap.InfoLevel
	}
	InitLogger(logPath, logInterface.GetLogMaxSize(), logInterface.GetLogBackups(), logInterface.GetLogMaxAge(), level, logInterface.GetLogInConsole())
	return nil
}

// InitLog 日志初始化
func InitLog(log LogInterface) error {
	if log != nil {
		logInterface = log
	}
	return reloadLog()
}

func init() {
	logInterface = defaultLogger{}
}
