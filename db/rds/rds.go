package tyrds

import (
	"errors"
	redigo "github.com/garyburd/redigo/redis"
	"time"
)

// var pool *redigo.Pool

type RdsServer struct {
	host string
	conn redigo.Conn
	pool *redigo.Pool
}

func init() {

}

func newPool(server string, password string, db int) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: time.Duration(240) * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", server, redigo.DialDatabase(db), redigo.DialPassword(password))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func ConnectRedis(host string) *RdsServer {
	if len(host) == 0 {
		return nil
	}

	// conn, err := redigo.Dial("tcp", host)
	// if err != nil {
	// 	fmt.Printf("Connect to redis error", err)
	// 	return nil
	// }

	return &RdsServer{
		host: host,
		// conn: conn,
		pool: newPool(host, "123456", 0),
	}
}

func (rs *RdsServer) GetConn() redigo.Conn {
	rs.conn = rs.pool.Get()
	return rs.conn
}

func (rs *RdsServer) SET(key, value string) int {
	if len(key) == 0 {
		return -1
	}

	if len(value) == 0 {
		return -1
	}

	reply, err := rs.conn.Do("SET", key, value)
	if err != nil {
		return -2
	}

	switch reply.(type) {
	case string:
	case int:
	}

	if _, ok := reply.(string); ok {
	}

	return 0
}

func (rs *RdsServer) SETEX(key, value string, ex int) int {
	if len(key) == 0 {
		return -1
	}

	if len(value) == 0 {
		return -1
	}
	reply, err := rs.conn.Do("SETEX", key, ex, value)
	if err != nil {
		return -2
	}

	switch reply.(type) {
	case string:
	case int:
	}

	return 0
}

func (rs *RdsServer) GET(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key is empty")
	}

	reply, er := redigo.String(rs.conn.Do("GET", key))
	if er != nil {
		return "", er
	}

	return reply, nil
}

func (rs *RdsServer) EXPIRE(key string, t int64) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("key is empty")
	}

	reply, er := rs.conn.Do("EXPIRE", "EXPIRE:"+key, t)
	if er != nil {
		return false, er
	}

	if reply == int64(1) {
	}

	return true, nil
}

func (rs *RdsServer) INCR(key string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("key is empty")
	}

	reply, er := rs.conn.Do("INCR", key)
	if er != nil {
		return false, er
	}

	if reply == int64(1) {
	}

	return true, nil
}

func (rs *RdsServer) DECR(key string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("key is empty")
	}

	reply, er := rs.conn.Do("DECR", key)
	if er != nil {
		return false, er
	}

	if reply == int64(1) {
	}

	return true, nil
}

func (rs *RdsServer) EXISTS(key string) (bool, error) {
	if len(key) == 0 {
		return false, errors.New("key is empty")
	}

	reply, er := rs.conn.Do("EXISTS", key)
	if er != nil {
		return false, er
	}

	if reply.(int64) == 1 {
	}

	return true, nil
}
