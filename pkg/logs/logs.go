package logs

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
	"time"
)

type Config struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAges    int
	Compress   bool
	Level      string
}

var once sync.Once
var zapLog *zap.Logger

// Init 初始化Logger
func Init(cfg Config) {
	once.Do(func() {
		writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAges)
		encoder := getEncoderConfig()
		var l = new(zapcore.Level)
		if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
			panic(err)
		}

		var core zapcore.Core
		if gin.Mode() == gin.DebugMode {
			core = zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.Lock(os.Stdout)),
				zapcore.DebugLevel,
			)
		} else {
			core = zapcore.NewCore(encoder, writeSyncer, l)
		}
		filed := zap.Fields(zap.String("serviceName", "blog"))

		zapLog = zap.New(core, zap.AddCaller(), filed)
		zap.ReplaceGlobals(zapLog)
	})

}

//设置日志格式
func getEncoderConfig() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

//设置日志切割
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Sync() {
	zapLog.Sync()
}
