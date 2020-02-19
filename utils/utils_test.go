package tyutils

import (
	"log"
	"testing"
)

func TestCompareMapStringString(t *testing.T) {
	m1 := map[string]string{
		"a": "aaaa",
		"b": "bbbb1",
	}

	m2 := map[string]string{
		"a": "aaaa",
		"b": "bbbb",
	}

	ok := CompareMapStringString(m1, m2)
	if ok {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestRandomIntger(t *testing.T) {
	v := RandomIntger(4)
	if v > 0 {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestRandomString(t *testing.T) {
	s := RandomString(32)
	if s != "" {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestGenerateSID(t *testing.T) {
	s := GenerateSID()
	if s != "" {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestValidatePhone(t *testing.T) {
	ok := ValidatePhone("136327678991")
	if ok {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestValidateEmail(t *testing.T) {
	ok := ValidateEmail("828292929@qq.com")
	if ok {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}

func TestHashCode(t *testing.T) {
	log.Println("HashCode:", HashCode(UUID()))
}
