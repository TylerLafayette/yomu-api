package drivers

import (
	"github.com/TylerLafayette/yomu-api/cache"
	"github.com/go-redis/redis/v7"
)

// RedisDriver is a cache driver meant to work with the Redis database.
type RedisDriver struct {
	client *redis.Client
}

// NewRedisDriver creates and returns a *RedisDriver.
func NewRedisDriver() *RedisDriver {
	return &RedisDriver{}
}

// Init initializes a Redis connection.
func (d *RedisDriver) Init(address string, password string, db int) error {
	// Create a new Redis client.
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	// Send a ping to make sure the server is responding.
	_, err := client.Ping().Result()
	if err != nil {
		// Ping failed, return error.
		return err
	}

	// Set the RedisDriver client to the Redis client.
	d.client = client
	return nil
}

// Get gets a value from the Redis cache.
func (d *RedisDriver) Get(key cache.Key) (cache.Value, error) {
	// Attempt to retrieve the value from Redis.
	result, err := d.client.Get(string(key)).Result()
	if err != nil {
		// Get failed, return error.
		return nil, err
	}

	// Return the result.
	return cache.Value(result), nil
}

// Set sets a value in the Redis cache.
func (d *RedisDriver) Set(key cache.Key, value cache.Value) error {
	return d.client.Set(string(key), value, 0).Err()
}

// Delete deletes a value in the Redis cache.
func (d *RedisDriver) Delete(key cache.Key) error {
	return d.client.Del(string(key)).Err()
}
