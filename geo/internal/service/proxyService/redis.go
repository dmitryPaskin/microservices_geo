package proxyService

import (
	"encoding/json"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/go-redis/redis"
	"time"
)

type ClientRedis struct {
	*redis.Client
}

func NewRedis(address, password string) (*ClientRedis, error) {
	client := ClientRedis{
		redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       0,
		}),
	}

	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &client, nil
}

func (c *ClientRedis) CheckCacheInRedis(key string) (*[]*model.Address, error) {
	resultString, err := c.Client.Get("").Result()

	var result []*model.Address
	if err := json.Unmarshal([]byte(resultString), &result); err != nil {
		return nil, err
	}
	return &result, err
}

func (c *ClientRedis) SaveCacheInRedis(key string, result interface{}) error {
	if err := c.Client.Set(key, result, 10*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}
