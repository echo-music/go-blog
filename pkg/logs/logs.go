package logs

import (
	"github.com/echo-music/go-blog/pkg/known"
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
var writeSyncer zapcore.WriteSyncer

// Init 初始化Logger
func Init(cfg Config) {
	once.Do(func() {

		lumberJackLogger := &lumberjack.Logger{
			Filename:   cfg.FileName,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAges,
		}
		writeSyncer = zapcore.AddSync(lumberJackLogger)

		var l = new(zapcore.Level)
		if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
			panic(err)
		}

		var core zapcore.Core
		encoder := getEncoderConfig()
		if gin.Mode() == gin.DebugMode {
			core = zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.Lock(os.Stdout)),
				zapcore.DebugLevel,
			)
		} else {
			core = zapcore.NewCore(encoder, writeSyncer, l)
		}

		filed := zap.Fields(zap.String("serviceName", "go-blog"))
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

func Sync() {
	zapLog.Sync()
}

func Writer() zapcore.WriteSyncer {
	return writeSyncer
}

func Ctx(c *gin.Context) *zap.Logger {
	l := zapLog
	if requestId := c.Value(known.XRequestIDKey); requestId != nil {
		l = l.With(zap.Any(known.XRequestIDKey, requestId))
	}
	return l
}
