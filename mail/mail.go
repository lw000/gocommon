package tymail

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
)

// {
// "From": "2241172930@qq.com",
// "Pass": "ltdxvumslixddjea",
// "To": [
// "373102227@qq.com"
// ],
// "host": "smtp.qq.com",
// "port": 465
// }

type JsonConfig struct {
	From string   `json:"Form"`
	To   []string `json:"To"`
	Pass string   `json:"pass"`
	Host string   `json:"host"`
	Port int64    `json:"port"`
}

func LoadConfig(file string) (*JsonConfig, error) {
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
	return fmt.Sprintf("{From:%s To:%v Pass:%s Host:%s Port:%d}", c.From, c.To, c.Pass, c.Host, c.Port)
}

func SendMail(cfg *JsonConfig, subject string, body string) error {
	if cfg == nil {
		return errors.New("MailConfig is nil")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "tuyue"+"<"+cfg.From+">")
	// 收件人
	m.SetHeader("To", cfg.To...)
	// 抄送
	m.SetHeader("Cc", cfg.From)
	// 暗送
	// m.SetHeader("BCc",  cfg.From)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("这是附件")
	d := gomail.NewDialer(cfg.Host, int(cfg.Port), cfg.From, cfg.Pass)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
