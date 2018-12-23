package infrastructure

import (
	"github.com/go-redis/redis"
	"encoding/json"
	"fmt"
	"errors"
)

type Redis interface {
	Get(key string, db int) ([]byte, error)
	Set(key string, val []byte, db int) error
	GetDB(keyRedisDB string, serviceName string) (int, error)
	Ping() error
}

type RedisType struct {
	config Config
}

func NewRedis(config Config) RedisType {
	return RedisType{
		config: config,
	}
}

func (r RedisType) Ping() error {
	client, err := r.newClient(1)
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
}

func (r RedisType) newClient(db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r RedisType) Get(key string, db int) ([]byte, error) {

	client, err := r.newClient(db)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	if key == "*" {
		return r.getAll(client)
	}

	return r.get(key, client)
}

func (r RedisType) get(key string, client *redis.Client) ([]byte, error) {
	blob, err := client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (r RedisType) getAll(client *redis.Client) ([]byte, error) {

	blob, err := client.Get("*").Bytes()
	if err != nil {
		return nil, err
	}

	var keys []string
	if err = json.Unmarshal(blob, &keys); err != nil {
		return nil, err
	}

	res := map[string][]byte{}
	for _, key := range keys {
		val, err := r.get(key, client)
		if err != nil {
			return nil, err
		}
		res[key] = val
	}

	blob, err = json.Marshal(res)
	if err != nil {
		return nil, err
	}

	return blob, nil
}

func (r RedisType) Set(key string, value []byte, db int) error {
	client, err := r.newClient(db)
	if err != nil {
		return err
	}
	defer client.Close()

	return r.set(key, value, client)
}

func (r RedisType) set(key string, value []byte, client *redis.Client) error {

	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisType) GetDB(keyRedisDB string, serviceName string) (int, error) {
	if keyRedisDB == "" {
		return 0, errors.New(fmt.Sprintf("Key -> %s dont exist in gin Params", serviceName))
	}

	db, ok := r.config.RedisDB[keyRedisDB]
	if !ok {
		return 0, errors.New(fmt.Sprintf("Key -> %s dont exist in RedisDB config", keyRedisDB))
	}

	return db, nil
}
