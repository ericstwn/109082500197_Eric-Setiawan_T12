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

	total := 0
	for i := 0; i < n; i++ {
		total += angka[i]
	}

	rata := float64(total) / float64(n)

	fmt.Println(rata)
}
