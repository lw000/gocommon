package main

import (
	"gocommon/test"

	log "github.com/alecthomas/log4go"
)

func main() {
	log.LoadConfiguration("conf/log4go.xml")
	test.Test()
}
