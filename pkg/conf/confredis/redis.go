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
	prefix      string
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

	if r.pool == nil {
		r.initial()
	}
}

func (r *Redis) initial() {
	addr := fmt.Sprintf("%s:%d", r.Host, r.Port)
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

func (r *Redis) get(key string) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	_key := r.Prefix(key)
	return conn.Do("GET", _key)
}

func (r *Redis) Get(key string) (reply interface{}, err error) {
	return r.get(key)
}

func (r *Redis) GetString(key string) (string, error) {
	return redis.String(r.get(key))
}

func (r *Redis) GetInt(key string) (int, error) {
	return redis.Int(r.get(key))
}

func (r *Redis) Set(key string, value interface{}) error {
	conn := r.pool.Get()
	defer conn.Close()

	_key := r.Prefix(key)
	_, err := conn.Do("SET", _key, value)
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

func (r *Redis) Prefix(key string) (outkey string) {
	if r.prefix != "" {
		return fmt.Sprintf("%s__%s", r.prefix, key)
	}
	return key
}
