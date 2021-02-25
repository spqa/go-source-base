package core

import (
	_ "embed"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
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
	"mcm-api/pkg/document"
	"mcm-api/pkg/faculty"
	"mcm-api/pkg/log"
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

//go:embed model.conf
var enforcerModel string

func ProvideEnforcer(db *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Logger.Panic("Init enforcer adapter failed", zap.Error(err))
	}
	m, err := model.NewModelFromString(enforcerModel)
	if err != nil {
		log.Logger.Panic("Init enforcer model failed", zap.Error(err))
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Logger.Panic("Init enforcer failed", zap.Error(err))
	}
	return e
}

var InfraSet = wire.NewSet(ProvideConfig, ProvideDB, ProvideRedis, ProvideEnforcer)
var HandlerSet = wire.NewSet(document.NewDocumentHandler, user.NewUserHandler, authz.NewAuthHandler, faculty.NewHandler)
