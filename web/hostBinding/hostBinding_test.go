package hostBinding

import (
	"reflect"
	"testing"
)

func TestDomainBinding_ErrText(t *testing.T) {
	type fields struct {
		domains map[string]bool
		errText string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "TestDomainBinding_SetErrText", fields: fields{domains: map[string]bool{"www.tuyue.com": true}, errText: "禁止访问"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &hostBinding{
				domains: tt.fields.domains,
				errText: tt.fields.errText,
			}
			if got := d.ErrText(); got != tt.want {
				t.Errorf("DomainBinding.ErrText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainBinding_SetErrText(t *testing.T) {
	type fields struct {
		domains map[string]bool
		errText string
	}
	type args struct {
		errText string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "TestDomainBinding_SetErrText", fields: fields{domains: map[string]bool{"www.tuyue.com": true}, errText: "禁止访问"}, args: args{"www.tuyue.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &hostBinding{
				domains: tt.fields.domains,
				errText: tt.fields.errText,
			}
			d.SetErrText(tt.args.errText)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *hostBinding
	}{
		{name: "TestNew", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainBinding_Binding(t *testing.T) {
	type fields struct {
		domains map[string]bool
		errText string
	}
	type args struct {
		domains []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "TestDomainBinding_Binding", fields: fields{domains: map[string]bool{"www.tuyue.com": true}, errText: ""}, args: args{[]string{"www.tuyue.com"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &hostBinding{
				domains: tt.fields.domains,
				errText: tt.fields.errText,
			}
			d.Binding(tt.args.domains...)
		})
	}
}

func TestDomainBinding_Allow(t *testing.T) {
	type fields struct {
		domains map[string]bool
		errText string
	}
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "TestDomainBinding_Allow", fields: fields{domains: map[string]bool{"www.tuyue.com": true}, errText: ""}, args: args{host: "www.tuyue.com"}, wantErr: false},
		{name: "TestDomainBinding_Allow", fields: fields{domains: map[string]bool{"www.tuyue.com:9090": true}, errText: ""}, args: args{host: "www.tuyue.com"}, wantErr: false},
		{name: "TestDomainBinding_Allow", fields: fields{domains: map[string]bool{"127.0.0.1": true}, errText: ""}, args: args{host: "127.0.0.1"}, wantErr: false},
		{name: "TestDomainBinding_Allow", fields: fields{domains: map[string]bool{"127.0.0.1:9090": true}, errText: ""}, args: args{host: "127.0.0.1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &hostBinding{
				domains: tt.fields.domains,
				errText: tt.fields.errText,
			}
			if err := d.Allow(tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("DomainBinding.Allow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
