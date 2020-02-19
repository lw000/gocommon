package tywscfg

import (
	"reflect"
	"testing"
)

func TestLoadJsonConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *JsonConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "json配置读取加载", args: struct{ file string }{file: string("conf.json")}, want: &JsonConfig{Host: "47.96.230.81:8830", Path: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadJsonConfig(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadJsonConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadJsonConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadIniConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    *IniConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "ini配置读取加载", args: struct{ file string }{file: string("conf.ini")}, want: &IniConfig{Host: "47.96.230.81:8830", Path: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadIniConfig(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadIniConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadIniConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
