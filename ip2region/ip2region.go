package tyip2region

import (
	"os"
	"sync"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

type IpRegionServer struct {
	sync.Mutex
	region *ip2region.Ip2Region
}

func NewIpRegionServer() *IpRegionServer {
	return &IpRegionServer{}
}

func (irs *IpRegionServer) LoadData(db string) error {
	_, err := os.Stat(db)
	if os.IsNotExist(err) {
		return err
	}

	irs.region, err = ip2region.New(db)
	if err != nil {
		return err
	}
	return nil
}

func (irs *IpRegionServer) Close() {
	irs.region.Close()
}

func (irs *IpRegionServer) ConverIp(command string) (string, error) {
	irs.Lock()
	defer irs.Unlock()

	ip, err := irs.region.BtreeSearch(command)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}
