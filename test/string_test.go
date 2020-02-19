package test

import "testing"

func comp(s string) bool {
	if s == "" {
		return true
	}
	return false
}

func comp1(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func BenchmarkTest(b *testing.B) {
	s := "111111"
	for i := 0; i < b.N; i++ {
		if comp(s) {

		}
	}
}

func BenchmarkTest2(b *testing.B) {
	s := "111111111111"
	for i := 0; i < b.N; i++ {
		if comp1(s) {

		}
	}
}
