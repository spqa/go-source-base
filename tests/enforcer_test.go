package tests

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mcm-api/config"
	"mcm-api/internal/core"
	"mcm-api/pkg/authz"
	"mcm-api/pkg/log"
	"mcm-api/pkg/user"
	"testing"
)

func ProvideDB(config *config.Config) *gorm.DB {
	m, err := migrate.New(
		"file://../migrations",
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

func TestEnforcer(t *testing.T) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	config := core.ProvideConfig()
	db := ProvideDB(config)
	enforcer := core.ProvideEnforcer(db)
	_ = enforcer.LoadPolicy()
	for _, v := range enforcer.GetPolicy() {
		enforcer.RemovePolicy(v)
	}
	_ = enforcer.SavePolicy()
	_, _ = enforcer.AddRoleForUser("1", fmt.Sprint(user.Student))
	_, _ = enforcer.AddPolicy(fmt.Sprint(user.Student), string(authz.Contribution), string(authz.Read))
	_ = enforcer.SavePolicy()
	enforce, _ := enforcer.Enforce("1", string(authz.Contribution), string(authz.Read))
	if enforce == false {
		t.Error("Test failed")
	}
}
