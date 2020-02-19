package tymysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/junhsieh/goexamples/fieldbinding/fieldbinding"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
	// "github.com/jinzhu/gorm"
)

// {
// "username": "root",
// "password": "Aabbcc123!@#",
// "host": "47.96.230.81:3306",
// "database": "mservice",
// "MaxOpenConns": 20,
// "MaxOdleConns": 5
// }

// 不建议使用Prepared Statement, 在Sharding模式下禁用
// golang中推荐使用:
// user:password@tcp(db.hostname:3306)/db_name?charset=utf8mb4,utf8&interpolateParams=true&parseTime=True&loc=UTC
// interpolateParams=true, 表示不直接使用prepared statement

type JsonConfig struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	MaxOpenConns int64  `json:"MaxOpenConns"`
	MaxOdleConns int64  `json:"MaxOdleConns"`
}

type YamlConfig struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Database     string `yaml:"database"`
	MaxOpenConns int64  `yaml:"maxOpenConns"`
	MaxOdleConns int64  `yaml:"maxOdleConns"`
}

func (c *YamlConfig) LoadData(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}

func (c *YamlConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return c.LoadData(data)
}

func (c YamlConfig) String() string {
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Username, c.Password, c.Host, c.Database, c.MaxOpenConns, c.MaxOdleConns)
}

func (c YamlConfig) JSON() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func LoadWithYamlData(data []byte) (*YamlConfig, error) {
	cfg := &YamlConfig{}
	if err := cfg.LoadData(data); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadWithYaml(file string) (*YamlConfig, error) {
	cfg := &YamlConfig{}
	if err := cfg.Load(file); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadJsonWithData(data []byte) (*JsonConfig, error) {
	cfg := &JsonConfig{}
	if err := cfg.LoadData(data); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadJsonConfig(file string) (*JsonConfig, error) {
	cfg := &JsonConfig{}
	if err := cfg.Load(file); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *JsonConfig) LoadData(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}

func (c *JsonConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return c.LoadData(data)
}

func (c JsonConfig) String() string {
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Username, c.Password, c.Host, c.Database, c.MaxOpenConns, c.MaxOdleConns)
}

type Mysql struct {
	db *sql.DB
}

func (m *Mysql) DB() *sql.DB {
	return m.db
}

func (m *Mysql) Begin() (*sql.Tx, error) {
	return m.db.Begin()
}

func (m *Mysql) OpenWithYamlConfig(cfg *YamlConfig) error {
	c := make(map[string]interface{})
	c["username"] = cfg.Username
	c["password"] = cfg.Password
	c["host"] = cfg.Host
	c["database"] = cfg.Database
	c["MaxOpenConns"] = cfg.MaxOpenConns
	c["MaxOdleConns"] = cfg.MaxOdleConns
	return m.open(c)
}

func (m *Mysql) OpenWithJsonConfig(cfg *JsonConfig) error {
	c := make(map[string]interface{})
	c["username"] = cfg.Username
	c["password"] = cfg.Password
	c["host"] = cfg.Host
	c["database"] = cfg.Database
	c["MaxOpenConns"] = cfg.MaxOpenConns
	c["MaxOdleConns"] = cfg.MaxOdleConns
	return m.open(c)
}

func (m *Mysql) open(cfg map[string]interface{}) error {
	var (
		username     string
		password     string
		host         string
		database     string
		maxOpenConns int64
		maxOdleConns int64
		v            interface{}
		ok           bool
	)

	if v, ok = cfg["database"]; ok {
		database = v.(string)
	} else {
		return errors.New("database is empty")
	}

	if v, ok = cfg["username"]; ok {
		username = v.(string)
	} else {
		return errors.New("username is empty")
	}

	if v, ok = cfg["password"]; ok {
		password = v.(string)
	} else {
		return errors.New("password is empty")
	}

	if v, ok = cfg["host"]; ok {
		host = v.(string)
	} else {
		return errors.New("host is empty")
	}

	if v, ok = cfg["MaxOpenConns"]; ok {
		maxOpenConns = v.(int64)
	} else {
		maxOpenConns = 10
	}

	if v, ok = cfg["MaxOdleConns"]; ok {
		maxOdleConns = v.(int64)
	} else {
		maxOdleConns = 0
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8&parseTime=true&loc=UTC", username, password, host, database)
	m.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	m.db.SetConnMaxLifetime(time.Hour * time.Duration(3))
	m.db.SetMaxOpenConns(int(maxOpenConns))
	m.db.SetMaxIdleConns(int(maxOdleConns))

	if err = m.db.Ping(); err != nil {
		return err
	}

	return nil
}

// Query ...
func (m *Mysql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	var (
		err error
		// stmt *sql.Stmt
	)
	// stmt, err = m.db.Prepare(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()

	var rows *sql.Rows
	if rows, err = m.db.Query(query, args...); err != nil {
		return nil, err
	}
	defer rows.Close()

	var cols []string
	if cols, err = rows.Columns(); err != nil {
		return nil, err
	}

	fb := fieldbinding.NewFieldBinding()
	fb.PutFields(cols)

	results := make([]map[string]interface{}, 0, 16)
	for rows.Next() {
		err = rows.Scan(fb.GetFieldPtrArr()...)
		if err != nil {
			return nil, err
		}
		results = append(results, fb.GetFieldArr())
	}

	if err = rows.Err(); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return results, nil
}

// QueryWithTimeout ...
func (m *Mysql) QueryWithTimeout(timeout int, query string, args ...interface{}) ([]map[string]interface{}, error) {
	var (
		err  error
		rows *sql.Rows
		ctx  context.Context
	)

	ctx, _ = context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	if rows, err = m.db.QueryContext(ctx, query, args...); err != nil {
		return nil, err
	}
	defer rows.Close()

	var cols []string
	if cols, err = rows.Columns(); err != nil {
		return nil, err
	}

	fb := fieldbinding.NewFieldBinding()
	fb.PutFields(cols)

	results := make([]map[string]interface{}, 0, 16)
	for rows.Next() {
		err = rows.Scan(fb.GetFieldPtrArr()...)
		if err != nil {
			return nil, err
		}
		results = append(results, fb.GetFieldArr())
	}

	if err = rows.Err(); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return results, nil
}

func (m *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

// ExecWithTimeout 执行语句，通过超时
func (m *Mysql) ExecWithTimeout(timeout int, query string, args ...interface{}) (sql.Result, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	execctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	return stmt.ExecContext(execctx, args...)
}

func (m *Mysql) Close() error {
	return m.db.Close()
}
