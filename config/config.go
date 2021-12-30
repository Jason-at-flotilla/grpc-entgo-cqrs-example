package config

import (
	"cqrs-grpc-test/pkg/constants"
	"cqrs-grpc-test/pkg/logger"
	"cqrs-grpc-test/pkg/postgres"
	"cqrs-grpc-test/pkg/redis"
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Reader microservice config path")
}

type Config struct {
	ServiceName     string           `mapstructure:"serviceName"`
	Logger          *logger.Config   `mapstructure:"logger"`
	GRPC_READ       GRPC             `mapstructure:"grpc_read"`
	GRPC_WRITE      GRPC             `mapstructure:"grpc_Write"`
	Postgresql      *postgres.Config `mapstructure:"postgres"`
	Redis           *redis.Config    `mapstructure:"redis"`
	ServiceSettings ServiceSettings  `mapstructure:"serviceSettings"`
}

type GRPC struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

type MongoCollections struct {
	Products string `mapstructure:"products"`
}

type ServiceSettings struct {
	RedisProductPrefixKey string `mapstructure:"redisProductPrefixKey"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	grpcPort := os.Getenv(constants.GrpcPort)
	if grpcPort != "" {
		cfg.GRPC_READ.Port = grpcPort
	}
	postgresHost := os.Getenv(constants.PostgresqlHost)
	if postgresHost != "" {
		cfg.Postgresql.Host = postgresHost
	}
	postgresPort := os.Getenv(constants.PostgresqlPort)
	if postgresPort != "" {
		cfg.Postgresql.Port = postgresPort
	}
	redisAddr := os.Getenv(constants.RedisAddr)
	if redisAddr != "" {
		cfg.Redis.Addr = redisAddr
	}

	return cfg, nil
}
