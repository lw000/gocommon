package tylimiter

import (
	gocache "github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Options struct {
	TM int // TM代表每TM秒允许访问一次
}

type Item struct {
	currentAccessTime int64
	nextAccessTime    int64
}

type IPRateLimiter struct {
	c       *gocache.Cache
	rate    time.Duration
	errText string
	sync.RWMutex
}

var (
	DefaultOption = &Options{TM: 1}
)

func New(opt *Options) *IPRateLimiter {
	if opt == nil {
		opt = DefaultOption
	}
	return &IPRateLimiter{
		rate: time.Second * time.Duration(opt.TM),
		c:    gocache.New(time.Second*time.Duration(opt.TM), time.Hour*6),
	}
}

func (p *IPRateLimiter) ErrText() string {
	p.RLock()
	defer p.RUnlock()
	return p.errText
}

func (p *IPRateLimiter) SetErrText(errText string) {
	p.Lock()
	defer p.Unlock()
	p.errText = errText
}

func (p *IPRateLimiter) push(ip string) error {
	t := time.Now()
	item := &Item{currentAccessTime: t.Unix() / 1e6, nextAccessTime: t.Add(p.rate+time.Duration(1)).UnixNano() / 1e6}
	p.Lock()
	defer p.Unlock()
	return p.c.Add(ip, item, p.rate)
}

func (p *IPRateLimiter) get(ip string) bool {
	p.RLock()
	defer p.RUnlock()
	// 缓存如果还在存在，则本次不允许访问，等待缓存过期
	_, found := p.c.Get(ip)
	return found
}

// 是否允许本次访问
func (p *IPRateLimiter) Allow(ip string) bool {
	found := p.get(ip)
	if found {
		return false
	}

	er := p.push(ip)
	if er != nil {
		return true
	}

	return true
}
