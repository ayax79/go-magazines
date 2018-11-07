package dao

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type (
	// MagazineDAO Interface for crud operations on magazine
	MagazineDAO interface {
		Get(uuid.UUID) (Magazine, error)
		Put(Magazine) error
	}

	// RedisMagazineDAO implements MagazineDAO using Redis as the store
	RedisMagazineDAO struct {
		client redis.Client
	}

	// RedisConfig ton configure Redis
	RedisConfig struct {
		Address  string
		Password string
		DB       int
	}
)

// NewRedisMagazineDAO constructs a new MagazineDAO instance backed by the redis implementation
func NewRedisMagazineDAO(config RedisConfig) *MagazineDAO {
	client := redis.NewClient(config.ToOptions)
	return &RedisMagazineDAO{client: &client}
}

// Get retrieves a Magazine from redis
func (dao *RedisMagazineDAO) Get(magazineID uuid.UUID) (Magazine, error) {
		

}

func (dao *RedisMagazineDAO) Put(magazine Magazine) error {
	_, err := dao.client.TxPipelined(func pipe redis.Pipeliner) error {
			return pipe
				.HSet(magazine, "TITLE", magazine.Title)
				.Hset(magazine, "ISSUE", magazine.Issue)
		})
		.Exec()
	return err
	
}
