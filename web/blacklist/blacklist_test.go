package tyblacklist

import (
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name string
		args args
		want BlackList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ips...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blackList_ErrMsg(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errMsg  string
		ips     map[string]bool
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
			wl := &blackList{
				RWMutex: tt.fields.RWMutex,
				errMsg:  tt.fields.errMsg,
				ips:     tt.fields.ips,
			}
			if got := wl.ErrMsg(); got != tt.want {
				t.Errorf("blackList.ErrMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blackList_SetErrMsg(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errMsg  string
		ips     map[string]bool
	}
	type args struct {
		errMsg string
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
			wl := &blackList{
				RWMutex: tt.fields.RWMutex,
				errMsg:  tt.fields.errMsg,
				ips:     tt.fields.ips,
			}
			wl.SetErrMsg(tt.args.errMsg)
		})
	}
}

func Test_blackList_SetIps(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errMsg  string
		ips     map[string]bool
	}
	type args struct {
		ips []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Test_blackList_SetIps", fields: fields{errMsg: "黑名单禁止访问", ips: make(map[string]bool)}, args: args{ips: []string{"192.168.1.73"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wl := &blackList{
				RWMutex: tt.fields.RWMutex,
				errMsg:  tt.fields.errMsg,
				ips:     tt.fields.ips,
			}
			wl.SetIps(tt.args.ips...)
		})
	}
}

func Test_blackList_GetIps(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errMsg  string
		ips     map[string]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wl := &blackList{
				RWMutex: tt.fields.RWMutex,
				errMsg:  tt.fields.errMsg,
				ips:     tt.fields.ips,
			}
			if got := wl.GetIps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("blackList.GetIps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_blackList_Deny(t *testing.T) {
	type fields struct {
		RWMutex sync.RWMutex
		errMsg  string
		ips     map[string]bool
	}
	type args struct {
		clientIP string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Test_blackList_SetIps", fields: fields{errMsg: "黑名单禁止访问", ips: make(map[string]bool)}, args: args{clientIP: "192.168.1.73"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wl := &blackList{
				RWMutex: tt.fields.RWMutex,
				errMsg:  tt.fields.errMsg,
				ips:     tt.fields.ips,
			}
			if err := wl.Deny(tt.args.clientIP); (err != nil) != tt.wantErr {
				t.Errorf("blackList.Deny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
