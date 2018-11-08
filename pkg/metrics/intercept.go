package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/stats"
	"net"
	"strconv"
	"time"
)

type Handler interface {
	Dialer(f func(string, time.Duration) (net.Conn, error)) func(string, time.Duration) (net.Conn, error)
	TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context
	HandleRPC(ctx context.Context, stat stats.RPCStats)
	TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context
	HandleConn(ctx context.Context, stat stats.ConnStats)
	SetMonitoring(monitoring *monitoring)
	GetMonitoring() *monitoring
}

type Interceptor struct {
	monitoring *monitoring
	trackPeers bool
}

// Peers- ref: https://godoc.org/google.golang.org/grpc/peer
type InterceptorOpts struct {
	TrackPeers bool
}

func RegisterInterceptor(s ServiceInfoProvider, i *Interceptor) (err error) {
	if i.trackPeers {
		return nil
	}

	infos := s.GetServiceInfo()
	for sn, info := range infos {
		for _, m := range info.Methods {
			t := streamType(m.IsClientStream, m.IsServerStream)

			for c := uint32(0); c <= 15; c++ {
				requestLabels := prometheus.Labels{
					"service": sn,
					"handler": m.Name,
					"code":    codes.Code(c).String(),
					"type":    t,
				}
				messageLabels := prometheus.Labels{
					"service": sn,
					"handler": m.Name,
				}

				// server
				if _, err = i.monitoring.server.errors.GetMetricWith(requestLabels); err != nil {
					return err
				}
				if _, err = i.monitoring.server.requestsTotal.GetMetricWith(requestLabels); err != nil {
					return err
				}
				if _, err = i.monitoring.server.requestDuration.GetMetricWith(requestLabels); err != nil {
					return err
				}
				if m.IsClientStream {
					if _, err = i.monitoring.server.messagesReceived.GetMetricWith(messageLabels); err != nil {
						return err
					}
				}
				if m.IsServerStream {
					if _, err = i.monitoring.server.messagesSend.GetMetricWith(messageLabels); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func NewInterceptor(opts InterceptorOpts) *Interceptor {
	return &Interceptor{
		monitoring: initMetrics(opts.TrackPeers),
		trackPeers: opts.TrackPeers,
	}
}

func (i *Interceptor) Dialer(f func(string, time.Duration) (net.Conn, error)) func(string, time.Duration) (net.Conn, error) {
	return func(addr string, timeout time.Duration) (net.Conn, error) {
		i.monitoring.dialer.WithLabelValues(addr).Inc()
		return f(addr, timeout)
	}
}

func (i *Interceptor) Describe(in chan<- *prometheus.Desc) {
	i.monitoring.dialer.Describe(in)
	i.monitoring.server.Describe(in)
	i.monitoring.client.Describe(in)
}

func (i *Interceptor) Collect(in chan<- prometheus.Metric) {
	i.monitoring.dialer.Collect(in)
	i.monitoring.server.Collect(in)
	i.monitoring.client.Collect(in)
}

func (i *Interceptor) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	service, method := split(info.FullMethodName)

	return context.WithValue(ctx, RPCKey, prometheus.Labels{
		"fail_fast": strconv.FormatBool(info.FailFast),
		"service":   service,
		"handler":   method,
	})
}

func (i *Interceptor) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	lab, _ := ctx.Value(RPCKey).(prometheus.Labels)

	switch in := stat.(type) {
	case *stats.Begin:
		if in.IsClient() {
			i.monitoring.client.requests.With(lab).Inc()
		} else {
			i.monitoring.server.requests.With(lab).Inc()
		}
	case *stats.End:
		if in.IsClient() {
			i.monitoring.client.requests.With(lab).Dec()
		} else {
			i.monitoring.server.requests.With(lab).Dec()
		}
	}
}

func (i *Interceptor) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return context.WithValue(ctx, ConnKey, prometheus.Labels{
		"remote_addr": info.RemoteAddr.String(),
		"local_addr":  info.LocalAddr.String(),
	})
}

// HandleConn implements stats Handler interface.
func (i *Interceptor) HandleConn(ctx context.Context, stat stats.ConnStats) {
	lab, _ := ctx.Value(ConnKey).(prometheus.Labels)

	switch in := stat.(type) {
	case *stats.ConnBegin:
		if in.IsClient() {
			i.monitoring.client.connections.With(lab).Inc()
		} else {
			i.monitoring.server.connections.With(lab).Inc()
		}
	case *stats.ConnEnd:
		if in.IsClient() {
			i.monitoring.client.connections.With(lab).Dec()
		} else {
			i.monitoring.server.connections.With(lab).Dec()
		}
	}
}

func (i *Interceptor) GetMonitoring() *monitoring {
	return i.monitoring
}

func (i *Interceptor) SetMonitoring(monitoring *monitoring) {
	i.monitoring = monitoring
}