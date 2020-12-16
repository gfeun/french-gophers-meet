package main

import (
	"time"
)

func main() {
	i := uint16(0)

	for {
		println("Hello tinygo", i)
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
