package main

import (
	"context"
	"fmt"
	"github.com/gofunct/cloudnative-engineer/pkg/metrics"
	pb "github.com/gofunct/cloudnative-engineer/protobuf"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net/http"
	"net"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()
	grpcMetrics = metrics.NewInterceptor(metrics.InterceptorOpts{})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(grpcMetrics)
}

// NOTE: Graceful shutdown is missing. Don't use this demo in your production setup.
func main() {
	// Listen an actual port.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	demoServer := newDemoServer()

	serverOps := []grpc.ServerOption{
		grpc.StatsHandler(grpcMetrics),
		grpc.UnaryInterceptor(unaryServerInterceptors(
			otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
			grpcMetrics.UnaryServer(),
		)),
	}

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 3001)}
	grpcServer := grpc.NewServer(serverOps...)

	metrics.RegisterInterceptor(grpcServer, grpcMetrics)
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// Register your service.
	pb.RegisterDemoServiceServer(grpcServer, demoServer)
	metrics.RegisterInterceptor(grpcServer, grpcMetrics)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	// Start your gRPC server.
	log.Fatal(grpcServer.Serve(lis))


}

func grpcDispatcherByPath(grpcServer *grpc.Server, other http.Handler, path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == path {
			other.ServeHTTP(w, r)
		} else {
			grpcServer.ServeHTTP(w, r)
		}
	})
}

func unaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		wrap := func(current grpc.UnaryServerInterceptor, next grpc.UnaryHandler) grpc.UnaryHandler {
			return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return current(currentCtx, currentReq, info, next)
			}
		}
		chain := handler
		for _, i := range interceptors {
			chain = wrap(i, chain)
		}
		return chain(ctx, req)
	}
}

// DemoServiceServer defines a Server.
type DemoServiceServer struct{}

func newDemoServer() *DemoServiceServer {
	return &DemoServiceServer{}
}

// SayHello implements a interface defined by protobuf.
func (s *DemoServiceServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello %s", request.Name)}, nil
}
