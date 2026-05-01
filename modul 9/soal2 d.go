package main

import "fmt"

func main() {
	var angka [20]int
	var n, x int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan angka %d: ", i)
		fmt.Scan(&angka[i])
	}

	fmt.Print("nilai x: ")
	fmt.Scan(&x)

	if x == 0 {
		return
	}

	fmt.Println("elemen dengan indeks kelipatan", x, ":")

	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Println(angka[i])
		}
	}
}
