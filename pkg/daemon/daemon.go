package daemon

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofunct/gofunct/pkg/tracing"
	pb "github.com/gofunct/gofunct/pkg/protobuf"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"strings"
	"sync"
	"testing"
	"time"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	mw "github.com/gofunct/gofunct/pkg/middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/gofunct/gofunct/pkg/metrics"
\)

const (
	Subsystem = "Daemon"
)

// DaemonOpts it is constructor argument that can be passed to
// the NewDaemon constructor function.
type DaemonOpts struct {
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
type Daemon struct {
	opts          *DaemonOpts
	done          chan struct{}
	serverOptions []grpc.ServerOption
	clientOptions []grpc.DialOption
	postgres      *sql.DB
	logger        *zap.Logger
	server        *grpc.Server
	rpcListener   net.Listener
	debugListener net.Listener
	tracerCloser  io.Closer

	lock sync.Mutex
}

// NewDaemon allocates new daemon instance using given options.
func NewDaemon(opts *DaemonOpts) (*Daemon, error) {
	d := &Daemon{
		done:          make(chan struct{}),
		opts:          opts,
		logger:        opts.Logger,
		serverOptions: opts.RPCOptions,
		rpcListener:   opts.RPCListener,
		debugListener: opts.DebugListener,
	}
	return d, nil
}

// TestDaemon returns address of fully started in-memory daemon and closer to close it.
func TestDaemon(t *testing.T, opts TestDaemonOpts) (net.Addr, io.Closer) {
	l, err := net.Listen("tcp", "127.0.0.1:0") // any available address
	if err != nil {
		t.Fatalf("mnemosyne daemon tcp listener setup error: %s", err.Error())
	}

	d, err := NewDaemon(&DaemonOpts{
		IsTest:            true,
		ClusterListenAddr: l.Addr().String(),
		Logger:            zap.L(),
		PostgresAddress:   opts.StoragePostgresAddress,
		PostgresTable:     "session",
		PostgresSchema:    "mnemosyne",
		RPCListener:       l,
	})
	if err != nil {
		t.Fatalf("mnemosyne daemon cannot be instantiated: %s", err.Error())
	}
	if err := d.Run(); err != nil {
		t.Fatalf("mnemosyne daemon start error: %s", err.Error())
	}

	return d.Addr(), d
}

// Run starts daemon and all services within.
func (d *Daemon) Run() (err error) {
	var (
		tracer opentracing.Tracer
	)

	interceptor := metrics.NewInterceptor(metrics.InterceptorOpts{})

	if d.opts.TracingAgentAddress != "" {
		if tracer, d.tracerCloser, err = tracing.InitJaeger(
			Subsystem,
			d.opts.TracingAgentAddress,
			d.logger.Named("tracer"),
		); err != nil {
			return
		}
		d.logger.Info("tracing enabled", zap.String("agent_address", d.opts.TracingAgentAddress))
	} else {
		tracer = opentracing.NoopTracer{}
	}

	d.clientOptions = []grpc.DialOption{
		// User agent is required for example to determine if incoming request is internal.
		grpc.WithUserAgent(fmt.Sprintf("%s:%s", Subsystem, d.opts.Version)),
		grpc.WithStatsHandler(interceptor),
		grpc.WithDialer(interceptor.Dialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		})),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptors(
			interceptor.UnaryClient(),
			otgrpc.OpenTracingClientInterceptor(tracer)),
		),
		grpc.WithStreamInterceptor(interceptor.StreamClient()),
	}
	d.serverOptions = []grpc.ServerOption{
		grpc.StatsHandler(interceptor),
		grpc.UnaryInterceptor(mw.UnaryServerInterceptors(
			otgrpc.OpenTracingServerInterceptor(tracer),
			mw.ErrorInterceptor(d.logger),
			interceptor.UnaryServer(),
		)),
	}
	if d.opts.TLS {
		servCreds, err := credentials.NewServerTLSFromFile(d.opts.TLSCertFile, d.opts.TLSKeyFile)
		if err != nil {
			return err
		}
		d.serverOptions = append(d.serverOptions, grpc.Creds(servCreds))

		clientCreds, err := credentials.NewClientTLSFromFile(d.opts.TLSCertFile, "")
		if err != nil {
			return err
		}
		d.clientOptions = append(d.clientOptions, grpc.WithTransportCredentials(clientCreds))
	} else {
		d.clientOptions = append(d.clientOptions, grpc.WithInsecure())
	}

	d.server = grpc.NewServer(d.serverOptions...)


	grpc_health_v1.RegisterHealthServer(d.server, health.NewServer())

	if !d.opts.IsTest {

		prometheus.DefaultRegisterer.Register(interceptor)
		metrics.RegisterInterceptor(d.server, interceptor)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("localhost:%v", 9093),
		grpc.WithStatsHandler(interceptor),
		grpc.WithDialer(interceptor.Dialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		})),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptors(
			interceptor.UnaryClient(),
			otgrpc.OpenTracingClientInterceptor(tracer)),
		),
		grpc.WithStreamInterceptor(interceptor.StreamClient()),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		d.logger.Info("rpc server is running", zap.String("address", d.rpcListener.Addr().String()))

		if err := d.server.Serve(d.rpcListener); err != nil {
			if err == grpc.ErrServerStopped {
				d.logger.Info("grpc server has been stopped")
				return
			}

			if !strings.Contains(err.Error(), "use of closed network connection") {
				d.logger.Error("rpc server failure", zap.Error(err))
			}
		}
	}()

	if d.debugListener != nil {
		go func() {
			d.logger.Info("debug server is running", zap.String("address", d.debugListener.Addr().String()))

			mux := http.NewServeMux()
			mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
			mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
			mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
			mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))
			mux.Handle("/healthz", &livenessHandler{
				livenessResponse: livenessResponse{
					Version: d.opts.Version,
				},
				logger: d.logger,
			})
			mux.Handle("/healthr", &readinessHandler{
				livenessResponse: livenessResponse{
					Version: d.opts.Version,
				},
				logger:   d.logger,
			})
			if err := http.Serve(d.debugListener, mux); err != nil {
				d.logger.Error("debug server failure", zap.Error(err))
			}
		}()
	}

	return
}

// Close implements io.Closer interface.
func (d *Daemon) Close() (err error) {
	d.done <- struct{}{}
	d.server.GracefulStop()
	if d.postgres != nil {
		if err = d.postgres.Close(); err != nil {
			return
		}
	}
	if d.debugListener != nil {
		if err = d.debugListener.Close(); err != nil {
			return
		}
	}
	if d.tracerCloser != nil {
		if err = d.tracerCloser.Close(); err != nil {
			return
		}
	}
	return nil
}

// Addr returns net.Addr that rpc service is listening on.
func (d *Daemon) Addr() net.Addr {
	return d.rpcListener.Addr()
}

