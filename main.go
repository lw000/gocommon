package main

import (
	"github.com/lw000/gocommon/test"

	log "github.com/alecthomas/log4go"
)

func main() {
	log.LoadConfiguration("conf/log4go.xml")
	test.Test()
}
