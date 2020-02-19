package hostBinding

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
	"sync"
)

const (
	defaultErrorText = "access denied"
)

type HostBinding interface {
	SetErrText(errText string)
	ErrText() string
	Binding(domains ...string)
	Allow(host string) error
}

type hostBinding struct {
	sync.RWMutex
	hosts   map[string]bool
	errText string
}

func New() HostBinding {
	bind := &hostBinding{
		hosts:   make(map[string]bool),
		errText: defaultErrorText,
	}
	bind.init()
	return bind
}

func (bind *hostBinding) init() {
	bind.hosts["127.0.0.1"] = true
	bind.hosts["localhost"] = true
}

func (bind *hostBinding) ErrText() string {
	return bind.errText
}

func (bind *hostBinding) SetErrText(errText string) {
	bind.errText = errText
}

func (bind *hostBinding) Binding(hosts ...string) {
	bind.Lock()
	defer bind.Unlock()

	for _, host := range hosts {
		if host == "" {
			continue
		}

		bind.hosts[host] = true

		addrs, err := net.LookupHost(host)
		if err != nil {
			log.Error(err)
			break
		}

		for _, addr := range addrs {
			bind.hosts[addr] = true
		}
	}

	log.WithField("hosts", bind.hosts).Info("hosts")
}

func (bind *hostBinding) Allow(host string) error {
	if host == "" {
		return errors.New(bind.errText)
	}

	if len(bind.hosts) == 2 {
		return nil
	}

	hosts := strings.Split(host, ":")
	if len(hosts) >= 2 {
		var err error
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			log.Error(err)
			return errors.New(bind.errText)
		}
	} else {
		host = hosts[0]
	}

	_, exists := bind.hosts[host]
	if !exists {
		return errors.New(bind.errText)
	}

	return nil
}
