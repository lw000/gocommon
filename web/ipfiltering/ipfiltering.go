package ipfiltering

import (
	"errors"
	log "github.com/sirupsen/logrus"
	tyip2region "gocommon/ip2region"
	"strings"
)

/*
	IP地址过滤库，过滤掉部分IP不允许访问服务，可设置例外，类似于黑名单里面的白名单功能，例外的IP地址允许访问
*/

const (
	defaultErrorText = "国外IP，禁止访问"
)

type IPfiltering struct {
	ipserv       *tyip2region.IpRegionServer // ip地址转换库
	abroadAccess bool                        // 是否允许国外IP访问
	errText      string                      // 错误提示
	whiteList    map[string]bool             // 特权例外
}

func New(abroadAccess bool, privilegedIP ...string) *IPfiltering {
	filter := &IPfiltering{
		abroadAccess: abroadAccess,
		errText:      defaultErrorText,
		whiteList:    make(map[string]bool),
	}
	filter.SetException(privilegedIP...)
	return filter
}

func (f *IPfiltering) SetException(excep ...string) {
	for _, e := range excep {
		f.whiteList[e] = true
	}
	f.whiteList["127.0.0.1"] = true
	f.whiteList["::1"] = true
}

func (f *IPfiltering) ErrText() string {
	return f.errText
}

func (f *IPfiltering) SetErrText(errText string) {
	f.errText = errText
}

func (f *IPfiltering) Load(db string) error {
	f.ipserv = tyip2region.NewIpRegionServer()
	if err := f.ipserv.LoadData(db); err != nil {
		return err
	}
	return nil
}

func (f *IPfiltering) abroadIP(s string) bool {
	v := strings.Split(s, "|")
	if len(v) >= 6 {
		if v[1] != "中国" {
			return true
		}
	}

	return false
}

func (f *IPfiltering) localIP(s string) bool {
	v := strings.Split(s, "|")
	if len(v) >= 6 {
		if v[4] == "内网IP" && v[5] == "内网IP" {
			return true
		}
	}

	return false
}

func (f *IPfiltering) Allow(clientIP string) (bool, error) {
	_, exists := f.whiteList[clientIP]
	if exists {
		log.WithFields(log.Fields{"clientIP": clientIP}).Info("IP白名单，允许访问")
		return true, nil
	}

	s, err := f.ipserv.ConverIp(clientIP)
	if err != nil {
		return false, err
	}

	log.WithFields(log.Fields{"clientIP": clientIP}).Info(s)

	// 本机IP访问
	local := f.localIP(s)
	if local {
		log.WithFields(log.Fields{"clientIP": clientIP}).Info("内网IP访问")
		return true, nil
	}

	// 允许国外IP访问
	if !f.abroadAccess {
		ok := f.abroadIP(s)
		if ok {
			return false, errors.New(f.errText)
		}
	}

	return true, nil
}
