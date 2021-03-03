package config

import "fmt"

type Config struct {
	RedisAddr        string `mapstructure:"redis_addr"`
	RedisPassword    string `mapstructure:"redis_password"`
	RedisDb          int    `mapstructure:"redis_db"`
	S3MediaBucket    string `mapstructure:"s3_media_bucket"`
	DatabaseHost     string `mapstructure:"database_host"`
	DatabasePort     string `mapstructure:"database_port"`
	DatabaseUsername string `mapstructure:"database_username"`
	DatabasePassword string `mapstructure:"database_password"`
	DatabaseName     string `mapstructure:"database_name"`
	WebAppUrl        string `mapstructure:"web_app_url"`
	JwtSecret        string `mapstructure:"jwt_secret"`
	AdminEmail       string `mapstructure:"admin_email"`
	AdminPassword    string `mapstructure:"admin_password"`
	MediaBucket      string `mapstructure:"media_bucket"`
	ConverterService string `mapstructure:"converter_service"`
}

func (config *Config) GetDatabaseDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		config.DatabaseHost,
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
	)
}

func (config *Config) GetDatabaseUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseName,
	)
}
