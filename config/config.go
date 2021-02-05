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
