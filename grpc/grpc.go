package grpc

import (
	"context"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/khanhduyy/shopms-common/logger"
	"google.golang.org/grpc/codes"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	xgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Runner is interface to provide an implementation to run gRPC service
type Runner interface {
	Run(ctx context.Context) error
}

// StandardRunner running gRPC Server application
type StandardRunner struct {
	Address string
	Server  *xgrpc.Server
}

//Run gRPC server on given context.
func (std *StandardRunner) Run(ctx context.Context) error {
	var (
		addr = std.Address
	)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			logger.Errorf("network %v and address %v on error %v", addr, err)
		}
	}()
	s := std.Server
	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	logger.Infof("Listening on address %v", addr)
	return s.Serve(l)
}

func New(ctx context.Context) *xgrpc.Server {

	logEntry := logger.WithContext(ctx)
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(func(c codes.Code) log.Level {
			return grpc_logrus.DefaultClientCodeToLevel(c)
		}),
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	s := xgrpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_ctxtags.UnaryServerInterceptor(), grpc_logrus.UnaryServerInterceptor(logEntry, opts...)),
		grpc_middleware.WithStreamServerChain(grpc_ctxtags.StreamServerInterceptor(), grpc_logrus.StreamServerInterceptor(logEntry, opts...)),
	)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	return s
}
