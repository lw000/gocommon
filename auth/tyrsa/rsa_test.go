package tyrsa

import (
	"crypto"
	"crypto/rsa"
	"reflect"
	"testing"
)

func TestNewRSA(t *testing.T) {
	type args struct {
		publicKey  []byte
		privateKey []byte
	}
	tests := []struct {
		name string
		args args
		want *RSAUtil
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRSA(tt.args.publicKey, tt.args.privateKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRSA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_pubKey(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    *rsa.PublicKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.pubKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.pubKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAUtil.pubKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_priKey(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		want    *rsa.PrivateKey
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.priKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.priKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAUtil.priKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_EncryptByPublicKey(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		originData []byte
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
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.EncryptByPublicKey(tt.args.originData)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.EncryptByPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAUtil.EncryptByPublicKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_DecryptByPrivateKey(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		originData []byte
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
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.DecryptByPrivateKey(tt.args.originData)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.DecryptByPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAUtil.DecryptByPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_Sign(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		originData []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := r.Sign(tt.args.originData, crypto.MD5)
			if (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RSAUtil.Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAUtil_Verify(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		originData []byte
		signData   string
		hash       crypto.Hash
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
			r := &RSAUtil{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			if err := r.Verify(tt.args.originData, tt.args.signData, tt.args.hash); (err != nil) != tt.wantErr {
				t.Errorf("RSAUtil.Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
