package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("the value is %v.\n", data)
	}
}
