package redis

import (
	"github.com/gomodule/redigo/redis"

	"bitbucket.org/alien_soft/api_getaway/storage/repo"
)

type redisRepo struct {
	rds *redis.Pool
}

// NewMailRepo ...
func NewRedisRepo(rds *redis.Pool) repo.InMemoryStorageI {
	return &redisRepo{rds: rds}
}

func (r *redisRepo) Set(key, value string) (err error) {
	conn := r.rds.Get()
	defer conn.Close()

	_, err = conn.Do("SET", key, value)
	return
}

func (r *redisRepo) SetWithTTl(key, value string, seconds int) (err error) {
	conn := r.rds.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, seconds, value)
	return
}

func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rds.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}
