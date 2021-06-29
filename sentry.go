package sentry

import (
	"context"
	"net"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	http2 "github.com/go-kratos/kratos/v2/transport/http"
)

const valuesKey = "sentry"

type Option func(*options)

type options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

func WithRepanic(repanic bool) Option {
	return func(opts *options) {
		opts.Repanic = repanic
	}
}

func WithWaitForDelivery(waitForDelivery bool) Option {
	return func(opts *options) {
		opts.WaitForDelivery = waitForDelivery
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.Timeout = timeout
	}
}

// Server returns a new server middleware for Sentry.
func Server(opts ...Option) middleware.Middleware {
	options := options{Repanic: true}
	for _, o := range opts {
		o(&options)
	}
	if options.Timeout == 0 {
		options.Timeout = 2 * time.Second
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			hub := sentry.GetHubFromContext(ctx)
			if hub == nil {
				hub = sentry.CurrentHub().Clone()
			}
			scope := hub.Scope()
			if tr, ok := transport.FromServerContext(ctx); ok {
				switch tr.Kind() {
				case transport.KindGRPC:
					gtr := tr.(*grpc.Transport)
					scope.SetContext("gRPC", map[string]interface{}{
						"endpoint":  gtr.Endpoint(),
						"operation": gtr.Operation(),
					})
					headers := make(map[string]interface{})
					for _, k := range gtr.Header().Keys() {
						headers[k] = gtr.Header().Get(k)
					}
					scope.SetContext("Headers", headers)
				case transport.KindHTTP:
					htr := tr.(*http2.Transport)
					r := htr.Request()
					scope.SetRequest(r)
				}
			}

			context.WithValue(ctx, valuesKey, hub)
			defer recoverWithSentry(options, hub, ctx, req)
			return handler(ctx, req)
		}
	}
}

func recoverWithSentry(opts options, hub *sentry.Hub, ctx context.Context, req interface{}) {
	if err := recover(); err != nil {
		if !isBrokenPipeError(err) {
			eventID := hub.RecoverWithContext(
				context.WithValue(ctx, sentry.RequestContextKey, req),
				err,
			)
			if eventID != nil && opts.WaitForDelivery {
				hub.Flush(opts.Timeout)
			}
		}
		if opts.Repanic {
			panic(err)
		}
	}
}

func isBrokenPipeError(err interface{}) bool {
	if netErr, ok := err.(*net.OpError); ok {
		if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(sysErr.Error()), "broken pipe") ||
				strings.Contains(strings.ToLower(sysErr.Error()), "connection reset by peer") {
				return true
			}
		}
	}
	return false
}

// GetHubFromContext retrieves attached *sentry.Hub instance from context.
func GetHubFromContext(ctx context.Context) *sentry.Hub {
	if hub, ok := ctx.Value(valuesKey).(*sentry.Hub); ok {
		return hub
	}
	return nil
}
