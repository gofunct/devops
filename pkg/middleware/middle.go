package middleware

import (
	"github.com/gofunct/gofunct/pkg/logz"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

func UnaryServerInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
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

func UnaryClientInterceptors(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		buildChain := func(current grpc.UnaryClientInterceptor, next grpc.UnaryInvoker) grpc.UnaryInvoker {
			return func(currentCtx context.Context, currentMethod string, currentReq, currentReply interface{}, currentCC *grpc.ClientConn, currentOpts ...grpc.CallOption) error {
				return current(currentCtx, currentMethod, currentReq, currentReply, currentCC, next, currentOpts...)
			}
		}
		chain := invoker
		for _, i := range interceptors {
			chain = buildChain(i, chain)
		}
		return chain(ctx, method, req, reply, cc, opts...)
	}
}


func LoggerBackground(ctx context.Context, log *zap.Logger, fields ...zapcore.Field) *zap.Logger {
	l := log.With(fields...)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if rid, ok := md["request_id"]; ok && len(rid) >= 1 {
			l = l.With(zap.String("request_id", rid[0]))
		}
	}
	return l
}

func GrpcDispatcherByPath(grpcServer *grpc.Server, other http.Handler, path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == path {
			other.ServeHTTP(w, r)
		} else {
			grpcServer.ServeHTTP(w, r)
		}
	})
}
func ErrorInterceptor(log *zap.Logger) func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
	{
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			now := time.Now()

			if md, ok := metadata.FromIncomingContext(ctx); ok {
				ctx = metadata.NewOutgoingContext(ctx, metadata.MD{
					"request_id":                     md["request_id"],
				})
			}

			res, err := handler(ctx, req)

			code := status.Code(err)
			if err != nil && code != codes.OK {
				if code == codes.Unknown {
					switch err {
					default:
						if pqerr, ok := err.(*pq.Error); ok {
							switch pqerr.Code {
							case pq.ErrorCode("57014"):
								code = codes.Canceled
							}
						} else {
							code = codes.Internal
						}
					}
				}
				loggerBackground(ctx, log).Error("request failure",
					zap.String("error", status.Convert(err).Message()),
					logz.Ctx(ctx, info, code),
				)

				switch err {
				default:
					return nil, status.Errorf(status.Code(err), "gofunct: %s", status.Convert(err).Message())
				}
			}

			loggerBackground(ctx, log).Debug("request handled successfully",
				logz.Ctx(ctx, info, codes.OK),
				zap.Duration("elapsed", time.Since(now)),
			)
			return res, err
		}
	}
}

func loggerBackground(ctx context.Context, log *zap.Logger, fields ...zapcore.Field) *zap.Logger {
	l := log.With(fields...)
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if rid, ok := md["request_id"]; ok && len(rid) >= 1 {
			l = l.With(zap.String("request_id", rid[0]))
		}
	}
	return l
}