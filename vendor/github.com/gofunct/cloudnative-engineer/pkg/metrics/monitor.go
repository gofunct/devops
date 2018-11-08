package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	"strings"
	"google.golang.org/grpc"
)

var (
	RPCKey  ctxKey = 1
	ConnKey ctxKey = 2
)

type ServiceInfoProvider interface {
	GetServiceInfo() map[string]grpc.ServiceInfo
}

type ctxKey int

type monitoring struct {
	dialer *prometheus.CounterVec
	server *monitor
	client *monitor
}

type monitor struct {
	connections      *prometheus.GaugeVec
	requests         *prometheus.GaugeVec
	requestsTotal    *prometheus.CounterVec
	requestDuration  *prometheus.HistogramVec
	messagesReceived *prometheus.CounterVec
	messagesSend     *prometheus.CounterVec
	errors           *prometheus.CounterVec
}

type Observe interface {
	Describe(in chan<- *prometheus.Desc)
	Collect(in chan<- prometheus.Metric)
}


func (m *monitor) Describe(in chan<- *prometheus.Desc) {
	m.connections.Describe(in)
	m.requests.Describe(in)
	m.requestDuration.Describe(in)
	m.requestsTotal.Describe(in)
	m.messagesReceived.Describe(in)
	m.messagesSend.Describe(in)
	m.errors.Describe(in)
}

func (m *monitor) Collect(in chan<- prometheus.Metric) {
	m.connections.Collect(in)
	m.requests.Collect(in)
	m.requestDuration.Collect(in)
	m.requestsTotal.Collect(in)
	m.messagesReceived.Collect(in)
	m.messagesSend.Collect(in)
	m.errors.Collect(in)
}

func initMetrics(trackPeers bool) *monitoring {
	dialer := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "reconnects_total",
			Help:      "Total number of reconnects made by client.",
		},
		[]string{"address"},
	)

	serverConnections := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "connections",
			Help:      "Number of currently opened server side connections.",
		},
		[]string{"remote_addr", "local_addr"},
	)
	serverRequests := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "requests",
			Help:      "Number of currently processed server side rpc requests.",
		},
		[]string{"fail_fast", "handler", "service"},
	)
	serverRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "requests_total",
			Help:      "Total number of RPC requests received by server.",
		},
		appendIt(trackPeers, []string{"service", "handler", "code", "type"}, "peer"),
	)
	serverReceivedMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "received_messages_total",
			Help:      "Total number of RPC messages received by server.",
		},
		appendIt(trackPeers, []string{"service", "handler"}, "peer"),
	)
	serverSendMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "send_messages_total",
			Help:      "Total number of RPC messages send by server.",
		},
		appendIt(trackPeers, []string{"service", "handler"}, "peer"),
	)
	serverRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "request_duration_seconds",
			Help:      "The RPC request latencies in seconds on server side.",
		},
		appendIt(trackPeers, []string{"service", "handler", "code", "type"}, "peer"),
	)
	serverErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "server",
			Name:      "errors_total",
			Help:      "Total number of errors that happen during RPC calles on server side.",
		},
		appendIt(trackPeers, []string{"service", "handler", "code", "type"}, "peer"),
	)

	clientConnections := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "connections",
			Help:      "Number of currently opened client side connections.",
		},
		[]string{"remote_addr", "local_addr"},
	)
	clientRequests := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "requests",
			Help:      "Number of currently processed client side rpc requests.",
		},
		[]string{"fail_fast", "handler", "service"},
	)
	clientRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "requests_total",
			Help:      "Total number of RPC requests made by client.",
		},
		[]string{"service", "handler", "code", "type"},
	)
	clientReceivedMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "received_messages_total",
			Help:      "Total number of RPC messages received.",
		},
		[]string{"service", "handler"},
	)
	clientSendMessages := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "send_messages_total",
			Help:      "Total number of RPC messages send.",
		},
		[]string{"service", "handler"},
	)
	clientRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "request_duration_seconds",
			Help:      "The RPC request latencies in seconds on client side.",
		},
		[]string{"service", "handler", "code", "type"},
	)
	clientErrors := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "grpc",
			Subsystem: "client",
			Name:      "errors_total",
			Help:      "Total number of errors that happen during RPC calls.",
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

func streamType(clientStream, serverStream bool) string {
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

func split(name string) (string, string) {
	if i := strings.LastIndex(name, "/"); i >= 0 {
		return name[1:i], name[i+1:]
	}
	return "unknown", "unknown"
}

func peerVal(ctx context.Context) string {
	v, ok := peer.FromContext(ctx)
	if !ok {
		return "none"
	}
	return v.Addr.String()
}

func appendIt(ok bool, arr []string, val string) []string {
	if !ok {
		return arr
	}
	return append(arr, val)
}


