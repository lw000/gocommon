package test

import (
	tyutils "gocommon/utils"
	"log"
	"testing"
)

func TestTExec(t *testing.T) {
	v, er := tyutils.TExec("", func() (i interface{}, e error) {
		return "ok", nil
	})
	if er != nil {
		log.Panic(er)
	}
	log.Println(v)
}
