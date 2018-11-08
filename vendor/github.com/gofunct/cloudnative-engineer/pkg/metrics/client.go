package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type monitoredClientStream struct {
	grpc.ClientStream
	labels  prometheus.Labels
	monitor *monitor
}

func (i *Interceptor) StreamClient() grpc.StreamClientInterceptor {
	monitor := i.monitoring.client

	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		start := time.Now()

		client, err := streamer(ctx, desc, cc, method, opts...)
		code := grpc.Code(err)
		service, method := split(method)
		labels := prometheus.Labels{
			"service": service,
			"handler": method,
			"code":    code.String(),
			"type":    streamType(desc.ClientStreams, desc.ServerStreams),
		}
		if err != nil && code != codes.OK {
			monitor.errors.With(labels).Add(1)
		}

		monitor.requestDuration.With(labels).Observe(time.Since(start).Seconds())
		monitor.requestsTotal.With(labels).Add(1)

		return &monitoredClientStream{ClientStream: client, monitor: monitor, labels: prometheus.Labels{
			"service": service,
			"handler": method,
		}}, err
	}
}

func (i *Interceptor) UnaryClient() grpc.UnaryClientInterceptor {
	monitor := i.monitoring.client

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)
		code := grpc.Code(err)
		service, method := split(method)
		labels := prometheus.Labels{
			"service": service,
			"handler": method,
			"code":    code.String(),
			"type":    "unary",
		}
		if err != nil && code != codes.OK {
			monitor.errors.With(labels).Add(1)
		}

		monitor.requestDuration.With(labels).Observe(time.Since(start).Seconds())
		monitor.requestsTotal.With(labels).Add(1)

		return err
	}
}

func (mcs *monitoredClientStream) SendMsg(m interface{}) error {
	err := mcs.ClientStream.SendMsg(m)
	if err == nil {
		mcs.monitor.messagesSend.With(mcs.labels).Inc()
	}
	return err
}

func (mcs *monitoredClientStream) RecvMsg(m interface{}) error {
	err := mcs.ClientStream.RecvMsg(m)
	if err == nil {
		mcs.monitor.messagesReceived.With(mcs.labels).Inc()
	}
	return err
}
