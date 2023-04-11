// 参考SigNoz文档
// https://signoz.io/docs/instrumentation/golang/
package otel

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTracer 初始化全局 trace provider
// exporter 将数据发送到 endpoint
// 返回 exporter 的 Shutdown 方法
func InitTracer(endpoint, service string) (func(context.Context) error, error) {
	ctx := context.Background()
	name := attribute.String("service.name", service)
	lang := attribute.String("library.language", "go")
	resOpt := resource.WithAttributes(name, lang)
	resource, err := resource.New(ctx, resOpt)
	if err != nil {
		logrus.Errorln("resource new error", err)
		return nil, err
	}

	insecure := otlptracegrpc.WithInsecure()
	epOpt := otlptracegrpc.WithEndpoint(endpoint)
	client := otlptracegrpc.NewClient(insecure, epOpt)
	// 创建exporter并启动
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		logrus.Errorln("otlptrace new error", err)
		return nil, err
	}

	sampler := trace.AlwaysSample()
	samOpt := trace.WithSampler(sampler)
	batcher := trace.WithBatcher(exporter)
	traceResOpt := trace.WithResource(resource)
	tp := trace.NewTracerProvider(samOpt, batcher, traceResOpt)
	// 配置全局 trace provider
	otel.SetTracerProvider(tp)
	return exporter.Shutdown, nil
}