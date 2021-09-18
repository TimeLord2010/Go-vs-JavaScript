package main

import (
	"fmt"
	"time"
)

func callatz(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return (n * 3) + 1
	}
}

func callatzSeries(n int) int {
	i := 0
	fmt.Printf("n = %d\n", n)
	for n != 1 {
		n = callatz(n)
		i = i + 1
	}
	fmt.Printf("iterations: %d\n", i)
	fmt.Println("------")
	return n
}

func main() {
	begin := time.Now()
	var a [10000]int
	for i := 0; i < 10000; i++ {
		n := callatzSeries(i + 1)
		a[i] = n
	}
	end := time.Now()
	fmt.Println("Array: ", a)
	fmt.Printf("Execution time: %dms", end.Sub(begin).Milliseconds())
}
