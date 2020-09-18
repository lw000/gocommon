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
		want *AesUtil
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestNewAes() = %v, want %v", got, tt.want)
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
			a := AesUtil{
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
			a := AesUtil{
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
