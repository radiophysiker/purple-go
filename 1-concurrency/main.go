package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := generateNumbers(10)
	squares := squareNumbers(numbers)

	for sq := range squares {
		fmt.Println(sq)
	}
}

func generateNumbers(n int) <-chan int {
	out := make(chan int)
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < n; i++ {
			out <- r.Intn(101)
		}
		close(out)
	}()
	return out
}

func squareNumbers(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range input {
			out <- num * num
		}
		close(out)
	}()
	return out
}
