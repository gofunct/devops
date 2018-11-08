package tracing

import (
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"io"
	"github.com/uber/jaeger-client-go/config"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaeger(service, node, agentAddress string, log *zap.Logger) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Tags: []opentracing.Tag{{
			Key:   constant.Subsystem + ".listen",
			Value: node,
		}},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentAddress,
		},
	}

	tracer, closer, err := cfg.New(service, config.Logger(zapjaeger.NewLogger(log)))
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, nil
}