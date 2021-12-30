package server

import (
	pb "cqrs-grpc-test/api/contactpb"
	"cqrs-grpc-test/models"
	"net"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func (s *server) newWriterGrpcServer(model *models.Models, server string, port string) (func() error, *grpc.Server, error) {

	l, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			s.im.Logger,
		),
		),
	)
	if server == "write" {
		writeService := &WriteService{
			Models: model,
		}
		pb.RegisterWriteContactServiceServer(grpcServer, writeService)
	}
	if server == "read" {
		readService := &ReadService{
			Models: model,
		}
		pb.RegisterReadContactServiceServer(grpcServer, readService)
	}
	go func() {
		s.log.Infof("gRPC server is listening on port: %s", port)
		s.log.Fatal(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}
