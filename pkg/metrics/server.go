package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type monitoredServerStream struct {
	grpc.ServerStream
	labels  prometheus.Labels
	monitor *monitor
}

func (mss *monitoredServerStream) SendMsg(m interface{}) error {
	err := mss.ServerStream.SendMsg(m)
	if err == nil {
		mss.monitor.messagesSend.With(mss.labels).Inc()
	}
	return err
}

func (mss *monitoredServerStream) RecvMsg(m interface{}) error {
	err := mss.ServerStream.RecvMsg(m)
	if err == nil {
		mss.monitor.messagesReceived.With(mss.labels).Inc()
	}
	return err
}

func (i *Interceptor) UnaryServer() grpc.UnaryServerInterceptor {
	monitor := i.monitoring.server

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		res, err := handler(ctx, req)
		code := grpc.Code(err)
		service, method := split(info.FullMethod)

		labels := prometheus.Labels{
			"service": service,
			"handler": method,
			"code":    code.String(),
			"type":    "unary",
		}
		if i.trackPeers {
			labels["peer"] = peerVal(ctx)
		}
		if err != nil && code != codes.OK {
			monitor.errors.With(labels).Add(1)
		}

		monitor.requestDuration.With(labels).Observe(time.Since(start).Seconds())
		monitor.requestsTotal.With(labels).Add(1)

		return res, err
	}
}

func (i *Interceptor) StreamServer() grpc.StreamServerInterceptor {
	monitor := i.monitoring.server

	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()

		service, method := split(info.FullMethod)
		streamLabels := prometheus.Labels{
			"service": service,
			"handler": method,
		}
		if i.trackPeers {
			if ss != nil {
				streamLabels["peer"] = peerVal(ss.Context())
			} else {
				// mostly for testing purposes
				streamLabels["peer"] = "nil-server-stream"
			}
		}
		err := handler(srv, &monitoredServerStream{ServerStream: ss, labels: streamLabels, monitor: monitor})
		code := grpc.Code(err)
		labels := prometheus.Labels{
			"service": service,
			"handler": method,
			"code":    code.String(),
			"type":    streamType(info.IsClientStream, info.IsServerStream),
		}
		if i.trackPeers {
			if ss != nil {
				labels["peer"] = peerVal(ss.Context())
			} else {
				// mostly for testing purposes
				labels["peer"] = "nil-server-stream"
			}
		}
		if err != nil && code != codes.OK {
			monitor.errors.With(labels).Add(1)
		}

		monitor.requestDuration.With(labels).Observe(time.Since(start).Seconds())
		monitor.requestsTotal.With(labels).Add(1)

		return err
	}
}
