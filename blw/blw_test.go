package tyblw

import (
	"testing"
)

func TestCoroutinePool(t *testing.T) {
	cos := New()
	err := cos.Start(1000)
	if err != nil {
		t.Error(err)
		return
	}

	cos.Submit(func() {
		t.Log("this is coroutine poll test")
	})
}
