package regionFiltering

import (
	"errors"
	log "github.com/sirupsen/logrus"
	tyip2region "gocommon/ip2region"
	"strings"
)

/*
	IP区域过滤
*/

const (
	defaultErrorText = "限制区域"
)

type RegionFiltering struct {
	errText string                      // 错误信息
	ipserv  *tyip2region.IpRegionServer // ip地址转换库
}

func New() *RegionFiltering {
	filter := &RegionFiltering{
		errText: defaultErrorText,
	}
	return filter
}

func (regi *RegionFiltering) ErrText() string {
	return regi.errText
}

func (regi *RegionFiltering) SetErrText(errText string) {
	regi.errText = errText
}

func (regi *RegionFiltering) Load(db string) error {
	regi.ipserv = tyip2region.NewIpRegionServer()
	if err := regi.ipserv.LoadData(db); err != nil {
		return err
	}
	return nil
}

func (regi *RegionFiltering) Allow(clientIP string, regions ...string) error {
	if len(regions) == 0 {
		return nil
	}

	ipContent, err := regi.ipserv.ConverIp(clientIP)
	if err != nil {
		log.Error(err)
		return err
	}
	log.WithFields(log.Fields{"region": ipContent}).Info("IP·区域")

	values := strings.Split(ipContent, "|")
	if len(values) >= 6 {
		city := values[4]

		// 内网IP，允许访问
		if city == "内网IP" {
			return nil
		}

		// 禁止区域，禁止访问
		for _, cy := range regions {
			if strings.HasPrefix(cy, city) {
				log.Error(regi.errText)
				return errors.New(regi.errText)
			}
		}
		return nil
	}

	return errors.New(regi.errText)
}
