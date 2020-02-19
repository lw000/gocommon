package test

import (
	"fmt"
	"github.com/lw000/gocommon/db/rdsex"
	"github.com/lw000/gocommon/utils"
	"log"
	"testing"
	"time"
)

var rediscfg *tyrdsex.JsonConfig

func TestRunTest2(t *testing.T) {
	var err error
	rediscfg, err = tyrdsex.LoadJsonConfig("../conf/redis.json")
	if err != nil {
		log.Panic(err)
	}
	log.Println(rediscfg)

	rds := &tyrdsex.RdsServer{}
	err = rds.OpenWithJsonConfig(rediscfg)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err = rds.Close()
	}()

	for i := 0; i < 10; i++ {
		_, _ = rds.Set(fmt.Sprintf("user:name%d", i), "levi", -1)
	}

	r, err := rds.Get("user:name3")
	if err != nil {
		log.Panic(err)
	}

	log.Println(r)

	for i := 0; i < 10; i++ {
		token := tyutils.UUID()
		_, _ = rds.Set("tokens:"+token, "1111", time.Second*time.Duration(300))
	}
}
