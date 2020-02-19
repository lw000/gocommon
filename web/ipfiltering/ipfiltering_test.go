package ipfiltering

import (
	"github.com/lw000/gocommon/ip2region"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		abroadAccess bool
		privilegedIP []string
	}
	tests := []struct {
		name string
		args args
		want *IPfiltering
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.abroadAccess, tt.args.privilegedIP...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPfiltering_SetException(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
	}
	type args struct {
		excep []string
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
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			f.SetException(tt.args.excep...)
		})
	}
}

func TestIPfiltering_ErrText(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
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
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			if got := f.ErrText(); got != tt.want {
				t.Errorf("IPfiltering.ErrText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPfiltering_SetErrText(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
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
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			f.SetErrText(tt.args.errText)
		})
	}
}

func TestIPfiltering_Load(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
	}
	type args struct {
		db string
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
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			if err := f.Load(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("IPfiltering.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIPfiltering_abroadIP(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			if got := f.abroadIP(tt.args.s); got != tt.want {
				t.Errorf("IPfiltering.abroadIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPfiltering_localIP(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			if got := f.localIP(tt.args.s); got != tt.want {
				t.Errorf("IPfiltering.localIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPfiltering_Allow(t *testing.T) {
	type fields struct {
		ipserv       *tyip2region.IpRegionServer
		AbroadAccess bool
		errText      string
		privilegedIP map[string]bool
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &IPfiltering{
				ipserv:       tt.fields.ipserv,
				AbroadAccess: tt.fields.AbroadAccess,
				errText:      tt.fields.errText,
				privilegedIP: tt.fields.privilegedIP,
			}
			got, err := f.Allow(tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("IPfiltering.Allow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IPfiltering.Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}
