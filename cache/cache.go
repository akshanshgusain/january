package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]interface{}

func (c *RedisCache) Has(str string) (bool, error) {
	k := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()

	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	ok, err := redis.Bool(conn.Do("EXISTS", k))
	if err != nil {
		return false, err
	}

	return ok, nil
}
