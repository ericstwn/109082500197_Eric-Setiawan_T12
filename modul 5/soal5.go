package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bilangan(1, n)
}

func bilangan(i int, n int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	bilangan(i+2, n)
}
