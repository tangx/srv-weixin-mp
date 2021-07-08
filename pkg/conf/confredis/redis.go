package confredis

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type Redis struct {
	Host        string `env:""`
	Password    string `env:""`
	Port        int16  `env:""`
	DB          int    `env:""`
	maxIdle     int
	idleTimeout time.Duration
	pool        *redis.Pool
}

func (r *Redis) SetDefaults() {
	if r.Host == "" {
		r.Host = "127.0.0.1"
	}
	if r.Port == 0 {
		r.Port = 6379
	}

	if r.maxIdle == 0 {
		r.maxIdle = 3
	}

	if r.idleTimeout == 0 {
		r.idleTimeout = 30
	}

}

func (r *Redis) Init() {
	r.SetDefaults()

	r.initial()
}

func (r *Redis) initial() {
	addr := fmt.Sprintf("%s:%d", r.Host, r.Port)
	if r.pool == nil {
		dial, err := redis.Dial("tcp",
			addr,
			redis.DialPassword(r.Password),
			redis.DialDatabase(r.DB),
		)
		if err != nil {
			log.Fatal(err)
		}

		r.pool = &redis.Pool{
			MaxIdle:     r.maxIdle,
			IdleTimeout: r.idleTimeout * time.Second,
			Dial: func() (redis.Conn, error) {
				return dial, nil
			},
		}
	}
}

func (r *Redis) get(key interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	return conn.Do("GET", key)
}

func (r *Redis) Get(key interface{}) (reply interface{}, err error) {
	return r.get(key)
}

func (r *Redis) GetString(key interface{}) (string, error) {
	return redis.String(r.get(key))
}

func (r *Redis) GetInt(key interface{}) (int, error) {
	return redis.Int(r.get(key))
}

func (r *Redis) Set(key, value interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("redis: set key(%v) failed: %v", key, err))

	}
	return nil
}

func (r *Redis) Conn() redis.Conn {
	return r.pool.Get()
}

func (r *Redis) Do(cmd string, args ...interface{}) (reply interface{}, err error) {
	conn := r.pool.Get()
	defer conn.Close()

	return conn.Do(cmd, args...)
}
