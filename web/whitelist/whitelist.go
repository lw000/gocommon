package tywhitelist

import (
	"errors"
	"sync"
)

const (
	defaultErrorText = "IP访问受限，拒接服务"
)

type WhiteList interface {
	SetErrMsg(errMsg string)
	ErrMsg() string

	SetWhiteList(ips ...string)
	WhiteLists() []string

	Allow(clientIp string) error
}

type whiteList struct {
	sync.RWMutex
	errMsg string
	whites map[string]bool
}

func New(ips ...string) WhiteList {
	wl := &whiteList{
		whites: make(map[string]bool),
		errMsg: defaultErrorText,
	}
	wl.init()
	wl.SetWhiteList(ips...)
	return wl
}

func (wl *whiteList) init() {
	wl.whites["127.0.0.1"] = true
	wl.whites["::1"] = true
}

func (wl *whiteList) ErrMsg() string {
	return wl.errMsg
}

func (wl *whiteList) SetErrMsg(errMsg string) {
	wl.errMsg = errMsg
}

func (wl *whiteList) SetWhiteList(ips ...string) {
	wl.Lock()
	defer wl.Unlock()

	for _, ip := range ips {
		wl.whites[ip] = true
	}
}

func (wl *whiteList) WhiteLists() []string {
	wl.RLock()
	defer wl.RUnlock()

	var ips []string
	for k, ok := range wl.whites {
		if ok {
			ips = append(ips, k)
		}
	}
	return ips
}

func (wl *whiteList) Allow(clientIP string) error {
	if len(wl.whites) == 2 {
		return nil
	}

	wl.RLock()
	defer wl.RUnlock()

	_, exists := wl.whites[clientIP]
	if exists {
		return nil
	}

	return errors.New(wl.errMsg)
}
