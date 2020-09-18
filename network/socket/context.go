package tysockt

import (
	"github.com/labstack/gommon/log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}
