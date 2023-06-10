package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := fib(n)
	fmt.Println(ans)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return (fib(n-1) + fib(n-2))
}
