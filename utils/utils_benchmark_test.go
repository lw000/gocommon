package tyutils

import (
	"testing"
)

func BenchmarkHashCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// log.Println("HashCode", HashCode(UUID()))
		HashCode(UUID())
	}
}

func BenchmarkRandomIntger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// log.Println("RandomIntger", RandomIntger(19))
		RandomIntger(19)
	}
}

func BenchmarkGenerateSID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// log.Println("GenerateSID", GenerateSID())
		GenerateSID()
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// log.Println("UUID", UUID())
		UUID()
	}
}

func BenchmarkGetInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// v, _ := GetInt64(int64(i))
		// log.Println("GetInt64", v)
		_, _ = GetInt64(int64(i))
	}
}

func BenchmarkValidatePhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := RandomString(11)
		// log.Println("ValidatePhone", ValidatePhone(v))
		ValidatePhone(v)
	}
}

func BenchmarkReverse(b *testing.B) {
	v := RandomString(128)
	for i := 0; i < b.N; i++ {
		// log.Println("Reverse", Reverse(v))
		Reverse(v)
	}
}
