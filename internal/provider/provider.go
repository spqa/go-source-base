package provider

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/pkg/log"
)

func ProvideConfig() *config.Config {
	var cfg *config.Config
	_ = viper.Unmarshal(&cfg)
	return cfg
}

func ProvideDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetDatabaseDsn()), &gorm.Config{})
	if err != nil {
		log.Logger.Panic("Connect to database failed", zap.Error(err))
	}
	return db
}

var Set = wire.NewSet(ProvideConfig, ProvideDB)
