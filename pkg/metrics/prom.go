package metrics

import (

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

)

const userAgentLabel = "user_agent"


var (
	tagRPCKey  ctxKey = 1
	tagConnKey ctxKey = 2
)

func initMonitoring(trackPeers bool, constLabels prometheus.Labels) *monitoring {
	dialer := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "reconnects_total",
			Help:        "Total number of reconnects made by client.",
			ConstLabels: constLabels,
		},
		[]string{"address"},
	)

	serverConnections := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "connections",
			Help:        "Number of currently opened server side connections.",
			ConstLabels: constLabels,
		},
		[]string{"remote_addr", "local_addr", userAgentLabel},
	)
	serverRequests := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "requests",
			Help:        "Number of currently processed server side rpc requests.",
			ConstLabels: constLabels,
		},
		[]string{"fail_fast", "handler", "service", userAgentLabel},
	)
	serverRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "requests_total",
			Help:        "Total number of RPC requests received by server.",
			ConstLabels: constLabels,
		},
		appendIf(trackPeers, []string{"service", "handler", "code", "type", userAgentLabel}, "peer"),
	)
	serverReceivedMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "received_messages_total",
			Help:        "Total number of RPC messages received by server.",
			ConstLabels: constLabels,
		},
		appendIf(trackPeers, []string{"service", "handler", userAgentLabel}, "peer"),
	)
	serverSendMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "send_messages_total",
			Help:        "Total number of RPC messages send by server.",
			ConstLabels: constLabels,
		},
		appendIf(trackPeers, []string{"service", "handler", userAgentLabel}, "peer"),
	)
	serverRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "request_duration_seconds",
			Help:        "The RPC request latencies in seconds on server side.",
			ConstLabels: constLabels,
		},
		appendIf(trackPeers, []string{"service", "handler", "code", "type", userAgentLabel}, "peer"),
	)
	serverErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "server",
			Name:        "errors_total",
			Help:        "Total number of errors that happen during RPC calles on server side.",
			ConstLabels: constLabels,
		},
		appendIf(trackPeers, []string{"service", "handler", "code", "type", userAgentLabel}, "peer"),
	)

	clientConnections := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "connections",
			Help:        "Number of currently opened client side connections.",
			ConstLabels: constLabels,
		},
		[]string{"remote_addr", "local_addr"},
	)
	clientRequests := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "requests",
			Help:        "Number of currently processed client side rpc requests.",
			ConstLabels: constLabels,
		},
		[]string{"fail_fast", "handler", "service"},
	)
	clientRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "requests_total",
			Help:        "Total number of RPC requests made by client.",
			ConstLabels: constLabels,
		},
		[]string{"service", "handler", "code", "type"},
	)
	clientReceivedMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "received_messages_total",
			Help:        "Total number of RPC messages received.",
			ConstLabels: constLabels,
		},
		[]string{"service", "handler"},
	)
	clientSendMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "send_messages_total",
			Help:        "Total number of RPC messages send.",
			ConstLabels: constLabels,
		},
		[]string{"service", "handler"},
	)
	clientRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "request_duration_seconds",
			Help:        "The RPC request latencies in seconds on client side.",
			ConstLabels: constLabels,
		},
		[]string{"service", "handler", "code", "type"},
	)
	clientErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "grpc",
			Subsystem:   "client",
			Name:        "errors_total",
			Help:        "Total number of errors that happen during RPC calls.",
			ConstLabels: constLabels,
		},
		[]string{"service", "handler", "code", "type"},
	)

	return &monitoring{
		dialer: dialer,
		server: &monitor{
			connections:      serverConnections,
			requests:         serverRequests,
			requestsTotal:    serverRequestsTotal,
			requestDuration:  serverRequestDuration,
			messagesReceived: serverReceivedMessages,
			messagesSend:     serverSendMessages,
			errors:           serverErrors,
		},
		client: &monitor{
			connections:      clientConnections,
			requests:         clientRequests,
			requestsTotal:    clientRequestsTotal,
			requestDuration:  clientRequestDuration,
			messagesReceived: clientReceivedMessages,
			messagesSend:     clientSendMessages,
			errors:           clientErrors,
		},
	}
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

type MonitoredClientStream struct {
	grpc.ClientStream
	labels  prometheus.Labels
	monitor *monitor
}


func handlerType(clientStream, serverStream bool) string {
	switch {
	case !clientStream && !serverStream:
		return "unary"
	case !clientStream && serverStream:
		return "server_stream"
	case clientStream && !serverStream:
		return "client_stream"
	default:
		return "bidirectional_stream"
	}
}



func userAgent(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ua, ok := md["user-agent"]; ok {
			return ua[0]
		}
	}
	return "not-set"
}

func withUserAgentLabel(ctx context.Context, lab prometheus.Labels) prometheus.Labels {
	lab[userAgentLabel] = userAgent(ctx)
	return lab
}

func peerValue(ctx context.Context) string {
	v, ok := peer.FromContext(ctx)
	if !ok {
		return "none"
	}
	return v.Addr.String()
}

func appendIf(ok bool, arr []string, val string) []string {
	if !ok {
		return arr
	}
	return append(arr, val)
}