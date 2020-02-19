package tyrdsex

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
	"time"
	// "gopkg.in/redis.v4"
	"github.com/go-redis/redis"
)

// {
// "host": "127.0.0.1:6379",
// "psd": "123456",
// "db": 0,
// "poolSize": 20,
// "minIdleConns": 5
// }

type JsonConfig struct {
	Host         string `json:"host"`
	Psd          string `json:"psd"`
	Db           int64  `json:"db"`
	PoolSize     int64  `json:"poolSize"`
	MinIdleConns int64  `json:"minIdleConns"`
}

type YamlConfig struct {
	Host         string `yaml:"host"`
	Psd          string `yaml:"psd"`
	Db           int64  `yaml:"db"`
	PoolSize     int64  `yaml:"poolSize"`
	MinIdleConns int64  `yaml:"minIdleConns"`
}

type RedisStore interface {
	Encode() map[string]interface{}
	Decode(data map[string]string) error
}

type RdsServer struct {
	client *redis.Client
}

func LoadConfigWithYaml(file string) (*YamlConfig, error) {
	cfg := &YamlConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *YamlConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c YamlConfig) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Host, c.Psd, c.Db, c.PoolSize, c.MinIdleConns)
}

func LoadJsonConfig(file string) (*JsonConfig, error) {
	cfg := &JsonConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *JsonConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c JsonConfig) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Host, c.Psd, c.Db, c.PoolSize, c.MinIdleConns)
}

func (r *RdsServer) Client() *redis.Client {
	return r.client
}

func (r *RdsServer) connect(host string, psd string, db int64, poolSize int64, minIdleConns int64) error {
	if r == nil {
		return errors.New("object is nil")
	}

	if host == "" {
		return errors.New("redis host is empty")
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     psd,
		DB:           int(db),
		PoolSize:     int(poolSize),
		MinIdleConns: int(minIdleConns),
		MaxConnAge:   time.Hour * time.Duration(2),
	})

	pong := r.client.Ping()
	if strings.ToUpper(pong.Val()) != "PONG" {
		return errors.New(pong.Err().Error())
	}

	return nil
}

func (r *RdsServer) OpenWithJsonConfig(cfg *JsonConfig) error {
	if cfg == nil {
		return errors.New("config is nil")
	}

	return r.connect(cfg.Host, cfg.Psd, cfg.Db, cfg.PoolSize, cfg.MinIdleConns)
}

func (r *RdsServer) OpenWithYamlConfig(cfg *YamlConfig) error {
	if cfg == nil {
		return errors.New("config is nil")
	}

	return r.connect(cfg.Host, cfg.Psd, cfg.Db, cfg.PoolSize, cfg.MinIdleConns)
}

func (r *RdsServer) Close() error {
	if r == nil {
		return errors.New("object is nil")
	}

	if err := r.client.Close(); err != nil {
		return err
	}

	return nil
}

func (r *RdsServer) Pipe() *redis.Pipeline {
	pipe := r.client.Pipeline().(*redis.Pipeline)
	return pipe
}

// func (r *RdsServer) Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return r.client.Scan(cursor, match, count).Result()
// }

func (r *RdsServer) ScanKeys(match string, count int64, f func(keys []string)) error {
	var (
		err    error
		cursor uint64 = 0
		keys   []string
	)

	for {
		keys, cursor, err = r.client.Scan(cursor, match, count).Result()
		if err != nil {
			break
		}

		if len(keys) > 0 {
			f(keys)
		}

		if cursor == 0 {
			break
		}
	}

	return err
}

func (r *RdsServer) Del(keys ...string) (int64, error) {
	return r.client.Del(keys...).Result()
}

func (r *RdsServer) Keys(key string) ([]string, error) {
	return r.client.Keys(key).Result()
}

func (r *RdsServer) Exists(key ...string) bool {
	exist, er := r.client.Exists(key...).Result()
	if er != nil {
		return false
	}
	return exist == 1
}

func (r *RdsServer) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.client.SetNX(key, value, expiration).Result()
}

func (r *RdsServer) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *RdsServer) Set(key string, v interface{}, expiration time.Duration) (string, error) {
	return r.client.Set(key, v, expiration).Result()
}

func (r *RdsServer) GetSet(key string, v interface{}) (string, error) {
	return r.client.GetSet(key, v).Result()
}

func (r *RdsServer) HGetAll(key string) (map[string]string, error) {
	return r.client.HGetAll(key).Result()
}

func (r *RdsServer) HMSet(key string, fields map[string]interface{}) (string, error) {
	return r.client.HMSet(key, fields).Result()
}

func (r *RdsServer) HMGet(key string, fields ...string) ([]interface{}, error) {
	return r.client.HMGet(key, fields...).Result()
}

// func (r *RdsServer) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return r.client.HScan(key, cursor, match, count).Result()
// }

func (r *RdsServer) HScanValues(key string, match string, count int64, f func(key string, value string)) error {
	var (
		err    error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, err = r.client.HScan(key, cursor, match, count).Result()
		if err != nil {
			break
		}

		if len(values) > 0 {
			for i := 0; i < len(values); i = i + 2 {
				f(values[i], values[i+1])
			}
		}

		if cursor == 0 {
			break
		}
	}

	return err
}

func (r *RdsServer) SAdd(key string, members ...interface{}) (int64, error) {
	return r.client.SAdd(key, members...).Result()
}

func (r *RdsServer) SMembers(key string) ([]string, error) {
	return r.client.SMembers(key).Result()
}

// func (r *RdsServer) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return r.client.SScan(key, cursor, match, count).Result()
// }

func (r *RdsServer) SScanValues(key string, match string, count int64, f func(values []string)) error {
	var (
		err    error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, err = r.client.SScan(key, cursor, match, count).Result()
		if err != nil {
			break
		}

		if len(values) > 0 {
			f(values)
		}

		if cursor == 0 {
			break
		}
	}
	return err
}

func (r *RdsServer) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.ZScan(key, cursor, match, count).Result()
}

func (r *RdsServer) ZScanValues(key string, match string, count int64, f func(values []string)) error {
	var (
		err    error
		cursor uint64 = 0
		values []string
	)
	for {
		values, cursor, err = r.client.ZScan(key, cursor, match, count).Result()
		if err != nil {
			break
		}

		if len(values) > 0 {
			f(values)
		}

		if cursor == 0 {
			break
		}
	}
	return err
}

func (r *RdsServer) IncrBy(key string, value int64) (int64, error) {
	return r.client.IncrBy(key, value).Result()
}

func (r *RdsServer) ListenKeyExpired(f func(msg *redis.Message)) {
	evnetKey := "notify-keyspace-events"
	rdscfg, err := r.client.ConfigGet(evnetKey).Result()
	if err != nil {
		return
	}

	v := rdscfg[1]
	if v.(string) == "" {
		var s string
		s, err = r.client.ConfigSet(evnetKey, "Ex").Result()
		if err != nil {
			return
		}
		log.Println(s)
	}

	Pubsub := r.client.Subscribe("__keyevent@0__:expired")
	for {
		var msg *redis.Message
		msg, err = Pubsub.ReceiveMessage()
		if err != nil {
			log.Println(err)
			return
		}

		f(msg)
	}
}
