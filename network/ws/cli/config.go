package tywscfg

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
)

// {
// 		"host": "47.96.230.81:8830",
// 		"path": ""
// }

type IniConfig struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type JsonConfig struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

func LoadJsonConfig(file string) (*JsonConfig, error) {
	cfg := &JsonConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func LoadIniConfig(file string) (*IniConfig, error) {
	cfg := &IniConfig{}
	err := cfg.Load(file)
	return cfg, err
}

func (c JsonConfig) String() string {
	return fmt.Sprintf("{Host:%s Path:%s}", c.Host, c.Path)
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

func (c IniConfig) String() string {
	return fmt.Sprintf("{Host:%s Path:%s}", c.Host, c.Path)
}

func (c *IniConfig) Load(file string) error {
	f, err := goconfig.LoadConfigFile(file)
	if err != nil {
		return fmt.Errorf("读取配置文件失败[%s]", file)
	}
	if err = c.readCfg(f); err != nil {
		return err
	}
	return nil
}

func (c *IniConfig) readCfg(f *goconfig.ConfigFile) error {
	var err error
	section := "ws"
	c.Host, err = f.GetValue(section, "host")
	if err != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "host", err.Error())
	}

	c.Path, err = f.GetValue(section, "path")
	if err != nil {
		return fmt.Errorf("无法获取键值(%s):%s", "path", err.Error())
	}

	return nil
}
