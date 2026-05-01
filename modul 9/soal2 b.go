package main

import "fmt"

func main() {
	var angka [20]int
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan angka %d: ", i)
		fmt.Scan(&angka[i])
	}

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Println(angka[i])
		}
	}
}
