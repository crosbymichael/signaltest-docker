package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	for sig := range c {
		fmt.Printf("got signal %s %d\n", sig, sig)
	}
}
