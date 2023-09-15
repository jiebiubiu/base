package trace

import (
	"fmt"
	"io"
	"time"

	"github.com/Ho-J/base/logs"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger() (opentracing.Tracer, io.Closer, error) {
	var cfg = jaegercfg.Configuration{
		ServiceName: "snoopy global trace", // 对其发起请求的的调用链，叫什么服务
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			CollectorEndpoint:   "http://127.0.0.1:14268/api/traces",
			LocalAgentHostPort:  "http://127.0.0.1:6831",
			BufferFlushInterval: 1 * time.Second,
		},
	}

	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(logs.ZapJaegerLog{}),
		// 设置最大 Tag 长度，根据情况设置
		jaegercfg.MaxTagValueLength(65535),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, err
}
