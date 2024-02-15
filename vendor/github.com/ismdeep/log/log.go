package log

import (
	"context"
	"fmt"
	"math"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TraceIDKeyType trace id key type
type TraceIDKeyType string

const (
	// TraceIDKey trace id key
	TraceIDKey TraceIDKeyType = "traceId"
)

type Logger struct {
	ZapLogger *zap.Logger
}

// DefaultLogger default logger
var DefaultLogger *Logger

// 初始化日志配置
func init() {
	Init("console://[stdout]?level=debug&time_encoder=rfc3339&trace_level=error")
}

func Init(dsn string) {
	DefaultLogger, _ = New(dsn)
}

func New(dsn string) (*Logger, error) {
	cfg, err := ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	var logLevel zapcore.Level
	switch cfg.Parameters.Level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	}
	var multiWriteSyncer []zapcore.WriteSyncer
	for _, output := range cfg.Outputs {
		switch output {
		case "stdout":
			multiWriteSyncer = append(multiWriteSyncer, zapcore.AddSync(os.Stdout))
		case "file":
			multiWriteSyncer = append(multiWriteSyncer, zapcore.AddSync(&lumberjack.Logger{
				Filename:   cfg.Parameters.FilePath,
				MaxSize:    1024,
				MaxAge:     1,
				MaxBackups: math.MaxInt64,
				Compress:   true,
			}))
		}
	}

	// 时间格式
	zapCfg := zap.NewProductionEncoderConfig()
	switch strings.ToLower(cfg.Parameters.TimeEncoder) {
	case "iso08601":
		zapCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	case "rfc3339":
		zapCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	case "rfc3339nano":
		zapCfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	case "epoch":
		zapCfg.EncodeTime = zapcore.EpochTimeEncoder
	case "epoch_millis":
		zapCfg.EncodeTime = zapcore.EpochMillisTimeEncoder
	case "epoch_nanos":
		zapCfg.EncodeTime = zapcore.EpochNanosTimeEncoder
	default:
		zapCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	}

	// Caller格式
	zapCfg.EncodeCaller = nil
	switch strings.ToLower(cfg.Parameters.CallerEncoder) {
	case "short":
		zapCfg.EncodeCaller = zapcore.ShortCallerEncoder
	case "full":
		zapCfg.EncodeCaller = zapcore.FullCallerEncoder
	}

	var encoder zapcore.Encoder
	switch cfg.Encoder {
	case "console":
		encoder = zapcore.NewConsoleEncoder(zapCfg)
	case "json":
		encoder = zapcore.NewJSONEncoder(zapCfg)
	default:
		panic(fmt.Errorf("unsupport log encoder: %v", cfg.Encoder))
	}

	// 配置信息
	var options []zap.Option
	options = append(options, zap.AddCaller()) // 添加调用者信息

	// 级别日志，打印堆栈
	switch cfg.Parameters.TraceLevel {
	case "error":
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	case "warn":
		options = append(options, zap.AddStacktrace(zap.WarnLevel))
	case "info":
		options = append(options, zap.AddStacktrace(zap.InfoLevel))
	case "debug":
		options = append(options, zap.AddStacktrace(zap.DebugLevel))
	}

	// 开启文件及行号
	options = append(options, zap.Development())

	// 初始化配置
	logger := zap.New(
		zapcore.NewCore(
			encoder, // json格式日志（ELK渲染收集）
			zapcore.NewMultiWriteSyncer(multiWriteSyncer...), // 打印到控制台和文件
			logLevel, // 日志级别
		),
		options...,
	)

	return &Logger{
		ZapLogger: logger,
	}, nil
}

// WithContext 从指定的context返回一个zap实例（关键方法）
func WithContext(ctx context.Context) *zap.Logger {
	return DefaultLogger.WithContext(ctx)
}

func (receiver *Logger) WithContext(ctx context.Context) *zap.Logger {
	if v := ctx.Value(TraceIDKey); v != "" {
		if s, ok := v.(string); ok {
			return receiver.ZapLogger.With(zap.String("traceId", s))
		}
	}

	if v := ctx.Value(string(TraceIDKey)); v != "" {
		if s, ok := v.(string); ok {
			return receiver.ZapLogger.With(zap.String("traceId", s))
		}
	}

	return receiver.ZapLogger
}

// NewTraceContext new context with a traceID
func NewTraceContext(traceID string) context.Context {
	return context.WithValue(context.Background(), TraceIDKey, traceID)
}
