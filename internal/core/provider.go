package core

import (
	_ "embed"
	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/authz"
	"mcm-api/pkg/contributesession"
	"mcm-api/pkg/contribution"
	"mcm-api/pkg/document"
	"mcm-api/pkg/faculty"
	"mcm-api/pkg/log"
	"mcm-api/pkg/media"
	"mcm-api/pkg/queue"
	"mcm-api/pkg/user"
)

func ProvideConfig() *config.Config {
	var cfg *config.Config
	_ = viper.Unmarshal(&cfg)
	return cfg
}

func ProvideDB(config *config.Config) *gorm.DB {
	m, err := migrate.New(
		"file://./migrations",
		config.GetDatabaseUrl())
	if err != nil {
		log.Logger.Panic("Connect to database failed", zap.Error(err))
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Logger.Panic("Migration failed", zap.Error(err))
	}
	db, err := gorm.Open(postgres.Open(config.GetDatabaseDsn()), &gorm.Config{})
	if err != nil {
		log.Logger.Panic("Connect to database failed", zap.Error(err))
	}
	return db
}

func ProvideRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword, // no password set
		DB:       config.RedisDb,       // use default DB
	})
	return rdb
}

var InfraSet = wire.NewSet(
	ProvideConfig,
	ProvideDB,
	ProvideRedis,
	queue.InitializeRedisQueue,
)
var HandlerSet = wire.NewSet(
	document.NewDocumentHandler,
	user.NewUserHandler,
	authz.NewAuthHandler,
	faculty.NewHandler,
	media.NewHandler,
	contributesession.NewHandler,
	contribution.NewHandler,
)
