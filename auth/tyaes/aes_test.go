package tyaes

import (
	"reflect"
	"testing"
)

func TestNewAes(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want *AesStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAes(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesStruct_cbcEncrypt(t *testing.T) {
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
			a := AesStruct{
				key: tt.fields.key,
			}
			got, err := a.cbcEncrypt(tt.args.data, tt.args.fillMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesStruct.cbcEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesStruct.cbcEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesStruct_cbcDecrypt(t *testing.T) {
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
			a := AesStruct{
				key: tt.fields.key,
			}
			got, err := a.cbcDecrypt(tt.args.data, tt.args.fillMode)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesStruct.cbcDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesStruct.cbcDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesEncrypt(t *testing.T) {
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
			got, err := AesEncrypt(tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesDecrypt(t *testing.T) {
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
			got, err := AesDecrypt(tt.args.key, tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesDecrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
