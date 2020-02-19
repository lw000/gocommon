package tysys

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterOnInterrupt(fn func(sign os.Signal)) {
	c := make(chan os.Signal, 1)
	sigs := []os.Signal{os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT}
	signal.Notify(c, sigs...)
	go func() {
		sign := <-c
		signal.Stop(c)
		fn(sign)
		close(c)
		os.Exit(0)
	}()
}
