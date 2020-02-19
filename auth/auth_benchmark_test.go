package tyauth

import (
	"github.com/lw000/gocommon/auth/tybase64"
	"github.com/lw000/gocommon/auth/tymd5"
	"github.com/lw000/gocommon/auth/tysha"
	"testing"
	"v/utils"
)

func BenchmarkHex(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		Hex([]byte(s))
	}
}

func BenchmarkMD5(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		_, _ = tymd5.MD5([]byte(s))
	}
}

func BenchmarkSha512(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		_ = tysha.Sha512([]byte(s))
	}
}

func BenchmarkSha256(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		_ = tysha.Sha256([]byte(s))
	}
}

func BenchmarkSha224(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		_ = tysha.Sha224([]byte(s))
	}
}

func BenchmarkSha1(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {

		_ = tysha.Sha1([]byte(s))
	}
}

func BenchmarkB64Encode(b *testing.B) {
	s := tyutils.RandomString(64)
	for i := 0; i < b.N; i++ {
		_ = tybase64.B64Encode([]byte(s))
	}
}
