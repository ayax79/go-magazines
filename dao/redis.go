package dao

import (
	"fmt"
	"log"

	"github.com/ayax79/go-magazines/model"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type (
	// RedisMagazineDAO implements MagazineDAO using Redis as the store
	RedisMagazineDAO struct {
		client *redis.Client
	}
)

// NewRedisMagazineDAO constructs a new MagazineDAO instance backed by the redis implementation
func NewRedisMagazineDAO(config *RedisConfig) (*RedisMagazineDAO, error) {
	client := redis.NewClient(config.ToOptions())
	err := client.Ping().Err()
	if err == nil {
		return &RedisMagazineDAO{client: client}, nil
	}
	log.Fatalf("Error instantiating Redis Client: %#v", err)
	return nil, err
}

// Get retrieves a Magazine from redis
func (dao *RedisMagazineDAO) Get(magazineID uuid.UUID) (*model.Magazine, error) {
	result, err := dao.client.HGetAll(magazineID.String()).Result()
	fmt.Printf("Get result for id %s :%s\n", result, err)
	if result != nil {
		title := result["TITLE"]
		issue := result["ISSUE"]
		magazine := model.NewMagazine(magazineID, title, issue)
		fmt.Printf("Returning magazine: %s\n", magazine)
		return magazine, err
	}
	return nil, err

}

// Put inserts or updates records
func (dao *RedisMagazineDAO) Put(magazine *model.Magazine) error {
	key := magazine.MagazineID.String()
	pipe := dao.client.TxPipeline()
	pipe.HSet(key, "TITLE", magazine.Title)
	pipe.HSet(key, "ISSUE", magazine.Issue)
	result, err := pipe.Exec()
	if result != nil {
		fmt.Printf("Put result: %s\n", result)
	}
	if err != nil {
		fmt.Printf("Put error: %s\n", err)
	}
	return err
}
