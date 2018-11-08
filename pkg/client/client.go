package client

import (
	"database/sql"
	"fmt"
	"github.com/gofunct/gofunct/pkg/metrics"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"io"
	"net"
	"net/http"
	"net/http/pprof"
	"strings"
	"sync"
	"time"
)

// DaemonOpts it is constructor argument that can be passed to
// the NewDaemon constructor function.
type ClientOpts struct {
	Version             string
	IsTest              bool
	SessionTTL          time.Duration
	SessionTTC          time.Duration
	TLS                 bool
	TLSCertFile         string
	TLSKeyFile          string
	Storage             string
	PostgresAddress     string
	PostgresTable       string
	PostgresSchema      string
	Logger              *zap.Logger
	RPCOptions          []grpc.ServerOption
	RPCListener         net.Listener
	DebugListener       net.Listener
	ClusterListenAddr   string
	ClusterSeeds        []string
	TracingAgentAddress string
}

// TestDaemonOpts set of options that are used with TestDaemon instance.
type TestDaemonOpts struct {
	StoragePostgresAddress string
}

// Daemon represents single daemon instance that can be run.
type Client struct {
	opts          *ClientOpts
	done          chan struct{}
	clientOptions []grpc.DialOption
	logger        *zap.Logger
	rpcListener   net.Listener
	debugListener net.Listener
	tracerCloser  io.Closer

	lock sync.Mutex
}

// Run starts daemon and all services within.
func (c *Client) Run() (err error) {
	var (
		tracer opentracing.Tracer
	)
	if cl, err = initCluster(c.logger, c.opts.ClusterListenAddr, c.opts.ClusterSeeds...); err != nil {
		return
	}
	if err = c.initStorage(c.logger, c.opts.PostgresTable, c.opts.PostgresSchema); err != nil {
		return
	}

	interceptor := metrics.NewInterceptor(metrics.InterceptorOpts{})

	if c.opts.TracingAgentAddress != "" {
		if tracer, c.tracerCloser, err = initJaeger(
			constant.Subsystem,
			c.opts.ClusterListenAddr,
			c.opts.TracingAgentAddress,
			c.logger.Named("tracer"),
		); err != nil {
			return
		}
		c.logger.Info("tracing enabled", zap.String("agent_address", c.opts.TracingAgentAddress))
	} else {
		tracer = opentracing.NoopTracer{}
	}

	c.clientOptions = []grpc.DialOption{
		// User agent is required for example to determine if incoming request is internal.
		grpc.WithUserAgent(fmt.Sprintf("%s:%s", constant.Subsystem, c.opts.Version)),
		grpc.WithStatsHandler(interceptor),
		grpc.WithDialer(interceptor.Dialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		})),
		grpc.WithUnaryInterceptor(unaryClientInterceptors(
			interceptor.UnaryClient(),
			otgrpc.OpenTracingClientInterceptor(tracer)),
		),
		grpc.WithStreamInterceptor(interceptor.StreamClient()),
	}
	c.serverOptions = []grpc.ServerOption{
		grpc.StatsHandler(interceptor),
		grpc.UnaryInterceptor(unaryServerInterceptors(
			otgrpc.OpenTracingServerInterceptor(tracer),
			errorInterceptor(c.logger),
			interceptor.UnaryServer(),
		)),
	}
	if c.opts.TLS {
		servCreds, err := credentials.NewServerTLSFromFile(c.opts.TLSCertFile, c.opts.TLSKeyFile)
		if err != nil {
			return err
		}
		c.serverOptions = append(c.serverOptions, grpc.Creds(servCreds))

		clientCreds, err := credentials.NewClientTLSFromFile(c.opts.TLSCertFile, "")
		if err != nil {
			return err
		}
		c.clientOptions = append(c.clientOptions, grpc.WithTransportCredentials(clientCreds))
	} else {
		c.clientOptions = append(c.clientOptions, grpc.WithInsecure())
	}

	c.server = grpc.NewServer(c.serverOptions...)

	cache := cache.New(5*time.Second, constant.Subsystem)
	mnemosyneServer, err := newSessionManager(sessionManagerOpts{
		addr:    c.opts.ClusterListenAddr,
		cluster: cl,
		logger:  c.logger,
		storage: c.storage,
		ttc:     c.opts.SessionTTC,
		cache:   cache,
		tracer:  tracer,
	})
	if err != nil {
		return err
	}

	mnemosynerpc.RegisterSessionManagerServer(c.server, mnemosyneServer)
	grpc_health_v1.RegisterHealthServer(c.server, health.NewServer())

	if !c.opts.IsTest {
		prometheus.DefaultRegisterer.Register(c.storage.(storage.InstrumentedStorage))
		prometheus.DefaultRegisterer.Register(cache)
		prometheus.DefaultRegisterer.Register(mnemosyneServer)
		prometheus.DefaultRegisterer.Register(interceptor)
		promgrpc.RegisterInterceptor(c.server, interceptor)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = cl.Connect(ctx, c.clientOptions...); err != nil {
		return err
	}

	go func() {
		c.logger.Info("rpc server is running", zap.String("address", c.rpcListener.Addr().String()))

		if err := c.server.Serve(c.rpcListener); err != nil {
			if err == grpc.ErrServerStopped {
				c.logger.Info("grpc server has been stopped")
				return
			}

			if !strings.Contains(err.Error(), "use of closed network connection") {
				c.logger.Error("rpc server failure", zap.Error(err))
			}
		}
	}()

	if c.debugListener != nil {
		go func() {
			c.logger.Info("debug server is running", zap.String("address", c.debugListener.Addr().String()))

			mux := http.NewServeMux()
			mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
			mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
			mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
			mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
			mux.Handle("/healthz", &livenessHandler{
				livenessResponse: livenessResponse{
					Version: c.opts.Version,
				},
				logger: c.logger,
			})
			mux.Handle("/healthr", &readinessHandler{
				livenessResponse: livenessResponse{
					Version: c.opts.Version,
				},
				logger:   c.logger,
				postgres: c.postgres,
				cluster:  cl,
			})
			if err := http.Serve(c.debugListener, mux); err != nil {
				c.logger.Error("debug server failure", zap.Error(err))
			}
		}()
	}

	go mnemosyneServer.cleanup(c.done)

	return
}
