package server

import (
	"context"
	"cqrs-grpc-test/config"
	"cqrs-grpc-test/internal/interceptors"
	"cqrs-grpc-test/internal/metrics"
	"cqrs-grpc-test/models"
	"cqrs-grpc-test/models/postgres"
	"cqrs-grpc-test/pkg/entclient"
	"cqrs-grpc-test/pkg/logger"
	"cqrs-grpc-test/pkg/redis"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator"
	"github.com/pkg/errors"
)

type server struct {
	log logger.Logger
	cfg *config.Config
	v   *validator.Validate

	im interceptors.InterceptorManager

	metrics *metrics.WriterServiceMetrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.im = interceptors.NewInterceptorManager(s.log)
	s.metrics = metrics.NewWriterServiceMetrics(s.cfg)

	// redis init
	redis.Init(s.cfg.Redis)
	redis.GetClient(int(redis.CACHE_LOGIN_SESSION.Number()))

	// ent GetClient
	dataSource := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		s.cfg.Postgresql.Host,
		s.cfg.Postgresql.User,
		s.cfg.Postgresql.Password,
		s.cfg.Postgresql.DBName,
	)
	entclient.Init(dataSource, "postgres")
	entclient, err := entclient.GetInstance()
	if err != nil {
		log.Fatal(err)
	}

	if err := entclient.Ent.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}

	models := GetModel(entclient)

	closeGrpcServer, grpcServer, err := s.newWriterGrpcServer(models)
	if err != nil {
		return errors.Wrap(err, "NewScmGrpcServer")
	}
	defer closeGrpcServer() // nolint: errcheck

	<-ctx.Done()
	grpcServer.GracefulStop()

	return nil
}

func GetModel(client *entclient.MyEnt) *models.Models {

	return &models.Models{
		ContactModel: &postgres.ContactModel{Client: client.Ent},
	}

}
