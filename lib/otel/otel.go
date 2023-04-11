// 参考SigNoz文档
// https://signoz.io/docs/instrumentation/golang/
package otel

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

var ServiceName string

// InitTracer 初始化全局 trace provider
// exporter 将数据发送到 endpoint
// 返回 exporter 的 Shutdown 方法，在程序结束时调用
// endpoint 无法连接时，初始化正常，发送数据时报错
func InitTracer(endpoint, service string) (func(context.Context) error, error) {
	ServiceName = service

	ctx := context.Background()
	name := attribute.String("service.name", ServiceName)
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

// GinMiddleware gin中间件，404时会发送数据
func GinMiddleware() gin.HandlerFunc {
	return otelgin.Middleware(ServiceName)
}
