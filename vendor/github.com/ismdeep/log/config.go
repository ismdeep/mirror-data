package log

import (
	"errors"
	"strings"
)

// Config config
type Config struct {
	Encoder    string   // 编码器： console, json
	Outputs    []string // 日志输出方式 stdout, file
	Parameters struct {
		Level         string // level 日志输出级别： e.g. debug, info, warn, error
		TraceLevel    string // trace_level 是否开启错误上下文追踪
		TimeEncoder   string // time_encoder 时间编码器(时间格式)
		CallerEncoder string // caller_encoder 调用者格式 e.g. short, full
		FilePath      string // file_path output包含file时需要此参数
	}
}

// ParseConfig parse config from dsn
// dsn e.g. console://[stdout,file]?level=debug&trace_level=error
func ParseConfig(dsn string) (*Config, error) {
	var result Config

	// 1. encoder
	index1 := strings.Index(dsn, "://")
	if index1 < 0 {
		return nil, errors.New("invalid dsn format")
	}
	encoder := dsn[:index1]
	switch encoder {
	case "console", "json":
		result.Encoder = encoder
	default:
		return nil, errors.New("unsupported encoder")
	}

	// 2. outputs
	dsn = dsn[index1+3:]
	index := strings.Index(dsn, "]")
	if index < 0 {
		return nil, errors.New("invalid dsn format")
	}
	s := dsn[:index+1]
	if s[0] != '[' || s[len(s)-1] != ']' {
		return nil, errors.New("invalid dsn format")
	}
	s = s[1 : len(s)-1]
	outputs := strings.Split(s, ",")
	for _, output := range outputs {
		switch output {
		case "stdout", "file":
			result.Outputs = append(result.Outputs, output)
		default:
			return nil, errors.New("invalid output pipe")
		}
	}

	// 3. extract parameters
	s = dsn[index+1:]
	if len(s) >= 1 {
		if s[0] != '?' {
			return nil, errors.New("invalid dsn format")
		}
		s = s[1:]
		items := strings.Split(s, "&")
		for _, item := range items {
			ts := strings.Split(item, "=")
			if len(ts) != 2 {
				return nil, errors.New("invalid dsn format")
			}
			switch ts[0] {
			case "level":
				result.Parameters.Level = ts[1]
			case "trace_level":
				result.Parameters.TraceLevel = ts[1]
			case "file_path":
				result.Parameters.FilePath = ts[1]
			case "time_encoder":
				result.Parameters.TimeEncoder = ts[1]
			case "caller_encoder":
				result.Parameters.CallerEncoder = ts[1]
			}
		}
	}

	return &result, nil
}
