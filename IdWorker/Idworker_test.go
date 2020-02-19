package tyIdWorker

import (
	"testing"
)

func TestIdWorker(t *testing.T) {
	idw := IdWorker{}
	err := idw.Start(1)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 100; i++ {
		t.Log("|", i, "|", idw.Id())
		t.Log("|", i, "|", idw.String())
	}
}
