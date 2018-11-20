package dao

import (
	"github.com/go-redis/redis"
)

type (
	// RedisConfig ton configure Redis
	RedisConfig struct {
		Address  string
		Password string
		DB       int
	}
)

// NewRedisConfig constructs a new configuration instance
func NewRedisConfig(address string, password string, db int) *RedisConfig {
	return &RedisConfig{address, password, db}
}

// ToOptions converts the RedisConfig object to a redis.Options struct
func (config *RedisConfig) ToOptions() *redis.Options {
	return &redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	}
}
