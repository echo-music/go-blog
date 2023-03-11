package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerZap "github.com/uber/jaeger-client-go/log/zap"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
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

var ZapLog *zap.Logger
var once sync.Once

// Init 初始化Logger
func Init(cfg Config) {

	once.Do(func() {
		writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAges)
		encoder := getEncoder()
		var l = new(zapcore.Level)
		if err := l.UnmarshalText([]byte(cfg.Level)); err != nil {
			panic(err)
		}

		var core zapcore.Core
		if gin.Mode() == gin.DebugMode {
			core = zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

		} else {
			core = zapcore.NewCore(encoder, writeSyncer, l)
		}

		ZapLog = zap.New(core, zap.AddCaller())
		zap.ReplaceGlobals(ZapLog)

		t, closer, err := NewGlobalTracer()
		if err != nil {
			panic(err)
		}
		defer closer.Close()
		opentracing.SetGlobalTracer(t)
	})
}

//设置日志格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = getEncodeTime
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
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

func NewGlobalTracer() (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jMetricsFactory := metrics.NullFactory

	// 初始化 Tracer 实例
	cfg.ServiceName = "go-blog-service"

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jaegerZap.NewLogger(ZapLog)),
		jaegercfg.Metrics(jMetricsFactory),
		// 设置最大 Tag 长度，根据情况设置
		jaegercfg.MaxTagValueLength(65535),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		panic(err)
	}

	return tracer, closer, err
}
