package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Test %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	test()
	test()

	time.Sleep(12 * time.Second)
}
