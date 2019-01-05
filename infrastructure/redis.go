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
	Delete(key string, db int) error
	GetDB(keyRedisDB string, serviceName string) (int, error)
	LLen(key string, db int) (int64, error)
	LPop(key string, db int) ([]byte, error)
	LPush(key string, val []byte, db int) error
	RPush(key string, val []byte, db int) error
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

func (r RedisType) getKeys(client *redis.Client) ([]string, error) {
	blob, err := client.Get("*").Bytes()
	if err != nil {
		return nil, err
	}

	var keys []string
	if err = json.Unmarshal(blob, &keys); err != nil {
		return nil, err
	}

	return keys, nil
}

func (r RedisType) getAll(client *redis.Client) ([]byte, error) {

	keys, err := r.getKeys(client)
	if err != nil {
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

	blob, err := json.Marshal(res)
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

func (r RedisType) Delete(key string, db int) error {
	client, err := r.newClient(db)
	if err != nil {
		return err
	}
	defer client.Close()

	if key == "*" {
		return r.deleteAll(client)
	}

	return r.delete(key, client)
}

func (r RedisType) deleteAll(client *redis.Client) error {

	keys, err := r.getKeys(client)
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := r.delete(key, client); err != nil {
			return err
		}
	}

	return nil
}

func (r RedisType) delete(key string, client *redis.Client) error {

	if err := client.Del(key).Err(); err != nil {
		return err
	}

	return nil
}

func (r RedisType) LLen(key string, db int) (int64, error) {
	client, err := r.newClient(db)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	llen, err := r.lLen(key, client)
	if err != nil {
		return 0, err
	}

	return llen, nil
}

func (r RedisType) lLen(key string, client *redis.Client) (int64, error) {

	llen, err := client.LLen(key).Result()
	if err != nil {
		return 0, err
	}

	return llen, nil
}

func (r RedisType) RPush(key string, val []byte, db int) error {
	client, err := r.newClient(db)
	if err != nil {
		return err
	}
	defer client.Close()

	if err := r.rPush(key, val, client); err != nil {
		return err
	}

	return nil
}

func (r RedisType) rPush(key string, val []byte, client *redis.Client) error {

	if err := client.RPush(key, val).Err(); err != nil {
		return err
	}

	return nil
}

func (r RedisType) LPush(key string, val []byte, db int) error {
	client, err := r.newClient(db)
	if err != nil {
		return err
	}
	defer client.Close()

	if err := r.lPush(key, val, client); err != nil {
		return err
	}

	return nil
}

func (r RedisType) lPush(key string, val []byte, client *redis.Client) error {

	if err := client.LPush(key, val).Err(); err != nil {
		return err
	}

	return nil
}

func (r RedisType) LPop(key string, db int) ([]byte, error) {
	client, err := r.newClient(db)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	val, err := r.lPop(key, client)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r RedisType) lPop(key string, client *redis.Client) ([]byte, error) {

	val, err := client.LPop(key).Bytes()
	if err != nil {
		return nil, err
	}

	return val, nil
}
