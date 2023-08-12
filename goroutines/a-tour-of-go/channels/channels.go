package main

import "fmt"

func sum(nums []int, channel chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	channel <- sum
}

func main() {
	nums := []int{4, 6, 7, 2, 9}

	channel := make(chan int)
	go sum(nums[:len(nums)/2], channel)
	go sum(nums[len(nums)/2:], channel)
	x, y := <-channel, <-channel
	fmt.Printf("Result: %d + %d = %d", x, y, x+y)
}
