package tymgo

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	// "gopkg.in/mgo.v2/bson"
	"time"
)

// {
// "address": "127.0.0.1:6379",
// username:"levi"
// "password": "123456",
// "db": "log",
// }

type JsonConfig struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Db       string   `json:"db"`
}

type YamlConfig struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Db       string   `yaml:"db"`
}

type Mongo struct {
	session *mgo.Session
}

func LoadYamlConfig(file string) (*YamlConfig, error) {
	cfg := &YamlConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *YamlConfig) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	er := yaml.Unmarshal(data, c)
	if er != nil {
		return er
	}

	return nil
}

func (c YamlConfig) String() string {
	return fmt.Sprintf("{Address:%v Username:%s Password:%s Db:%s}", c.Address, c.Username, c.Password, c.Db)
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

	if err = json.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}

func (c JsonConfig) String() string {
	return fmt.Sprintf("{Address:%v Username:%s Password:%s Db:%s}", c.Address, c.Username, c.Password, c.Db)
}

func (m *Mongo) Open(cfg *JsonConfig) error {
	dailInfo := &mgo.DialInfo{
		Addrs:     cfg.Address,
		Direct:    false,
		Timeout:   time.Second * 15,
		Database:  cfg.Db,
		Source:    "",
		Username:  "",
		Password:  "",
		PoolLimit: 1024,
	}
	var er error
	m.session, er = mgo.DialWithInfo(dailInfo)
	if er != nil {
		return er
	}
	// set mode
	m.session.SetMode(mgo.Monotonic, true)

	return nil
}

// func (m *Mongo)Session() *mgo.Session{
// 	return m.session.Copy()
// }

func (m *Mongo) Close() {
	m.session.Close()
}

func (m *Mongo) DB(db string) *mgo.Database {
	return m.session.DB(db)
}
