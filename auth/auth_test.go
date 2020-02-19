package tyauth

import (
	"log"
	"testing"
)

func TestHex(t *testing.T) {
	s := Hex([]byte("11111111111111111111"))
	log.Println(s)
}
