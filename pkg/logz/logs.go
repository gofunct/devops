package logz

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
)

const encoderName = "gofunct-stackdriver"

func init() {
	if err := zap.RegisterEncoder(encoderName, func(cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
		return &Encoder{
			Encoder: zapcore.NewJSONEncoder(cfg),
		}, nil
	}); err != nil {
		panic(err)
	}
}

type Opts struct {
	Environment string
	Version     string
	Level       string
}

// Init allocates new logger based on given options.
func Init(opts Opts) (logger *zap.Logger, err error) {
	var (
		cfg     zap.Config
		options []zap.Option
		lvl     zapcore.Level
	)
	switch opts.Environment {
	case "production":
		cfg = zap.NewProductionConfig()
	case "stackdriver":
		cfg = NewStackdriverConfig()
		options = append(options, zap.Fields(zap.Object("serviceContext", &ServiceContext{
			Service: "mnemosyned",
			Version: opts.Version,
		})))
	case "development":
		cfg = zap.NewDevelopmentConfig()
	default:
		cfg = zap.NewProductionConfig()
	}

	if err = lvl.Set(opts.Level); err != nil {
		return nil, err
	}
	cfg.Level.SetLevel(lvl)

	logger, err = cfg.Build(options...)
	if err != nil {
		return nil, err
	}
	logger.Info("logger has been initialized", zap.String("environment", opts.Environment))

	return logger, nil
}

func NewStackdriverEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "eventTime",
		LevelKey:      "severity",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			switch l {
			case zapcore.DebugLevel:
				enc.AppendString("DEBUG")
			case zapcore.InfoLevel:
				enc.AppendString("INFO")
			case zapcore.WarnLevel:
				enc.AppendString("WARNING")
			case zapcore.ErrorLevel:
				enc.AppendString("ERROR")
			case zapcore.DPanicLevel:
				enc.AppendString("CRITICAL")
			case zapcore.PanicLevel:
				enc.AppendString("ALERT")
			case zapcore.FatalLevel:
				enc.AppendString("EMERGENCY")
			}
		},
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// Ctx ...
func Ctx(ctx context.Context, info *grpc.UnaryServerInfo, code codes.Code) zapcore.Field {
	logCtx := Context{
		HTTPRequest: HTTPRequest{
			Method:             info.FullMethod,
			ResponseStatusCode: code.String(),
		},
	}
	if p, ok := peer.FromContext(ctx); ok {
		logCtx.HTTPRequest.RemoteIP = p.Addr.String()
	}

	return zap.Object("context", logCtx)
}

type Encoder struct {
	zapcore.Encoder
}

func (e *Encoder) Clone() zapcore.Encoder {
	return &Encoder{
		Encoder: e.Encoder.Clone(),
	}
}
func (e *Encoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	if ent.Caller.Defined {
		for i, f := range fields {
			if f.Key == "context" && f.Type == zapcore.ObjectMarshalerType {
				if ctx, ok := f.Interface.(Context); ok {
					fields[i] = zapcore.Field{Type: zapcore.SkipType}
					fields = append(fields, zap.Object("context", Context{
						HTTPRequest: ctx.HTTPRequest,
						User:        ctx.User,
						reportLocation: reportLocation{
							FilePath:   ent.Caller.File,
							LineNumber: ent.Caller.Line,
						},
					}))
					ent.Caller.Defined = false
				}
				break
			}
		}
	}

	if ent.Caller.Defined {
		fields = append(fields, zap.Object("context", Context{
			reportLocation: reportLocation{
				FilePath:   ent.Caller.File,
				LineNumber: ent.Caller.Line,
			},
		}))
		ent.Caller.Defined = false
	}

	if ent.Stack != "" {
		ent.Message = ent.Message + "\n" + ent.Stack
		ent.Stack = ""
	}

	return e.Encoder.EncodeEntry(ent, fields)
}

// NewStackdriverConfig ...
func NewStackdriverConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         encoderName,
		EncoderConfig:    NewStackdriverEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
}

type ServiceContext struct {
	Service string `json:"service"`
	Version string `json:"version"`
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (sc ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if sc.Service == "" {
		return errors.New("service name is mandatory")
	}
	enc.AddString("service", sc.Service)
	enc.AddString("version", sc.Version)

	return nil
}

var (
	emptyHTTPRequest    HTTPRequest
	emptyReportLocation reportLocation
)

type Context struct {
	HTTPRequest    HTTPRequest
	User           string
	reportLocation reportLocation
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (c Context) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if c.HTTPRequest != emptyHTTPRequest {
		enc.AddObject("httpRequest", c.HTTPRequest)
	}
	if c.User != "" {
		enc.AddString("user", c.User)
	}
	if c.reportLocation != emptyReportLocation {
		enc.AddObject("reportLocation", c.reportLocation)
	}

	return nil
}

type HTTPRequest struct {
	Method             string `json:"method"`
	URL                string `json:"url"`
	UserAgent          string `json:"userAgent"`
	Referrer           string `json:"referrer"`
	ResponseStatusCode string `json:"responseStatusCode"`
	RemoteIP           string `json:"remoteIp"`
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (hr HTTPRequest) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("method", hr.Method)
	enc.AddString("url", hr.URL)
	enc.AddString("userAgent", hr.UserAgent)
	enc.AddString("referrer", hr.Referrer)
	enc.AddString("responseStatusCode", hr.ResponseStatusCode)
	enc.AddString("remoteIp", hr.RemoteIP)

	return nil
}

// ReportLocation ...
type reportLocation struct {
	FilePath     string `json:"filePath"`
	LineNumber   int    `json:"lineNumber"`
	FunctionName string `json:"functionName"`
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (rl reportLocation) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("filePath", rl.FilePath)
	enc.AddInt("lineNumber", rl.LineNumber)
	enc.AddString("functionName", rl.FunctionName)

	return nil
}