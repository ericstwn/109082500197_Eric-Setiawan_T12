package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	test(n)
}

func test(n int) {
	if n == 0 {
		return
	} else {
		fmt.Print(n)
		test(n - 1)
		fmt.Print(n)
	}
}
