package tydes

import (
	"reflect"
	"testing"
)

func TestNewDes(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want *DesUtil
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDES(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesStruct_ecbEncrypt(t *testing.T) {
	type fields struct {
		key []byte
	}
	type args struct {
		data     []byte
		fillMode int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DesUtil{
				key: tt.fields.key,
			}
			got, err := d.ecbEncrypt(tt.args.data, tt.args.fillMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesStruct.ecbEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DesStruct.ecbEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesStruct_ecbDecrypt(t *testing.T) {
	type fields struct {
		key []byte
	}
	type args struct {
		data     []byte
		fillMode int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DesUtil{
				key: tt.fields.key,
			}
			got, err := d.ecbDecrypt(tt.args.data, tt.args.fillMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesStruct.ecbDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DesStruct.ecbDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesEcbEncrypt(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DesEcbEncrypt(tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesEcbEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DesEcbEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesEcbDecrypt(t *testing.T) {
	type args struct {
		key []byte
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DesEcbDecrypt(tt.args.key, tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesEcbDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DesEcbDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
