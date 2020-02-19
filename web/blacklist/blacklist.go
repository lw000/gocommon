package tyblacklist

import (
	"errors"
	"sync"
)

const (
	defaultErrorText = "黑名单，禁止访问"
)

type BlackList interface {
	SetErrMsg(errMsg string)
	ErrMsg() string

	SetBlackList(ips ...string)
	BlackLists() []string

	Deny(clientIP string) error
}

type blackList struct {
	sync.RWMutex
	errMsg string
	blacks map[string]bool
}

func New(ips ...string) BlackList {
	bl := &blackList{
		blacks: make(map[string]bool),
		errMsg: defaultErrorText,
	}
	bl.SetBlackList(ips...)
	return bl
}

func (bl *blackList) ErrMsg() string {
	return bl.errMsg
}

func (bl *blackList) SetErrMsg(errMsg string) {
	bl.errMsg = errMsg
}

func (bl *blackList) SetBlackList(ips ...string) {
	bl.Lock()
	defer bl.Unlock()

	for _, ip := range ips {
		bl.blacks[ip] = true
	}
}

func (bl *blackList) BlackLists() []string {
	bl.RLock()
	defer bl.RUnlock()

	var ips []string
	for ip, ok := range bl.blacks {
		if ok {
			ips = append(ips, ip)
		}
	}
	return ips
}

func (bl *blackList) Deny(clientIP string) error {
	if clientIP == "" {
		return errors.New(bl.errMsg)
	}

	if len(bl.blacks) == 0 {
		return nil
	}

	bl.RLock()
	defer bl.RUnlock()

	_, exists := bl.blacks[clientIP]
	if exists {
		return errors.New(bl.errMsg)
	}

	return nil
}
