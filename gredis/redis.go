package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/lyydhl/go-base-module/serialize"
	"time"
)

var RedisConn *redis.Pool

// 初始化
func Setup(maxIdle, maxActive int, idleTimeout time.Duration, host, password string) error {
	RedisConn = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

// 存储值，默认 7200 秒
// key 键 [字符串类型]
// data 值 [数据不需要处理，方面里面会自动json化]
func SetDefault(key string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, 7200)
	if err != nil {
		return err
	}

	return nil
}

// 存储值
// key 键 [字符串类型]
// data 值 [数据不需要处理，方面里面会自动json化]
// time 时间 [秒 如果为 -1 则永久存储]
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	if time != -1 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// 存储值 存储二进制
// key 键 [字符串类型]
// data 值 [数据不需要处理，方面里面会自动序列化]
// time 时间 [秒]
func SetToBytes(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := serialize.GetBytes(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	if time != -1 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// 是否存在
// key 键 [字符串类型]
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// 获取数据
// key 键 [字符串类型]
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 删除数据
// key 键 [字符串类型]
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// 删除数据
// key 键 [字符串类型]
func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
