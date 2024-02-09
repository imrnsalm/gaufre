package main

import (
	"fmt"
	"time"
)

func main() {
	size := 100
	garden := make(chan int, size)
	burrow := make(chan int, size)

	go gopher(garden, burrow, 500)
	go gopher(garden, burrow, 875)

	for vegetable := 1; vegetable < size; vegetable++ {
		garden <- vegetable
	}
	close(garden)

	for food := 1; food < size; food++ {
		fmt.Printf("Gophers have stolen %d vegetables from the garden\n", <-burrow)
	}
}

func gopher(garden <-chan int, burrow chan<- int, speed time.Duration) {
	for vegetable := range garden {
		burrow <- pocket(vegetable)
		dig(speed)
	}
}

func pocket(food int) int {
	return food + 1
}

func dig(speed time.Duration) {
	time.Sleep(time.Millisecond * speed)
}
