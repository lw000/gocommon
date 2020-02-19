package tylimiter

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	gocache "github.com/patrickmn/go-cache"
)

func TestIPLimiter_ErrText(t *testing.T) {
	type fields struct {
		count     int
		c         *gocache.Cache
		frequency time.Duration
		errText   string
		RWMutex   sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IPLimiter{
				count:     tt.fields.count,
				c:         tt.fields.c,
				frequency: tt.fields.frequency,
				errText:   tt.fields.errText,
				RWMutex:   tt.fields.RWMutex,
			}
			if got := p.ErrText(); got != tt.want {
				t.Errorf("IPLimiter.ErrText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPLimiter_SetErrText(t *testing.T) {
	type fields struct {
		count     int
		c         *gocache.Cache
		frequency time.Duration
		errText   string
		RWMutex   sync.RWMutex
	}
	type args struct {
		errText string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IPLimiter{
				count:     tt.fields.count,
				c:         tt.fields.c,
				frequency: tt.fields.frequency,
				errText:   tt.fields.errText,
				RWMutex:   tt.fields.RWMutex,
			}
			p.SetErrText(tt.args.errText)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opt *Options
	}
	tests := []struct {
		name string
		args args
		want *IPLimiter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPLimiter_addLimit(t *testing.T) {
	type fields struct {
		count     int
		c         *gocache.Cache
		frequency time.Duration
		errText   string
		RWMutex   sync.RWMutex
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IPLimiter{
				count:     tt.fields.count,
				c:         tt.fields.c,
				frequency: tt.fields.frequency,
				errText:   tt.fields.errText,
				RWMutex:   tt.fields.RWMutex,
			}
			if err := p.addLimit(tt.args.ip); (err != nil) != tt.wantErr {
				t.Errorf("IPLimiter.addLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPLimiter_allow(t *testing.T) {
	type fields struct {
		count     int
		c         *gocache.Cache
		frequency time.Duration
		errText   string
		RWMutex   sync.RWMutex
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Item
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &IPLimiter{
				count:     tt.fields.count,
				c:         tt.fields.c,
				frequency: tt.fields.frequency,
				errText:   tt.fields.errText,
				RWMutex:   tt.fields.RWMutex,
			}
			got, got1 := p.allow(tt.args.ip)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPLimiter.allow() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IPLimiter.allow() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGinIpLimitHandler(t *testing.T) {
	type args struct {
		lim *IPLimiter
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GinIpLimitHandler(tt.args.lim); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GinIpLimitHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_String(t *testing.T) {
	type fields struct {
		Tm     string
		Nexttm string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Item{
				Tm:     tt.fields.Tm,
				Nexttm: tt.fields.Nexttm,
			}
			if got := e.String(); got != tt.want {
				t.Errorf("Item.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
