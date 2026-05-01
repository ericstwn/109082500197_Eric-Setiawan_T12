package main

import "fmt"

func main() {
	var angka [10]int
	var n, cari int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan angka %d: ", i)
		fmt.Scan(&angka[i])
	}

	fmt.Scan(&cari)

	frekuensi := 0
	for i := 0; i < n; i++ {
		if angka[i] == cari {
			frekuensi++
		}
	}

	fmt.Println(frekuensi)
}
