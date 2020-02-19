package tywhitelist

import (
	"log"
	"testing"
)

func TestWhiteListSrv(t *testing.T) {
	white := New()
	white.SetWhiteList("127.0.0.1", "192.168.1.73")
	white.SetErrMsg("whitelist error context")
	log.Println(white.WhiteList())
	log.Println(white.ErrMsg())
}
