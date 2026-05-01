package main

import "fmt"

func main() {
	var angka [20]int
	var n, hapus int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan angka %d: ", i)
		fmt.Scan(&angka[i])
	}

	fmt.Print("indeks yg dihapus: ")
	fmt.Scan(&hapus)

	for i := hapus; i < n-1; i++ {
		angka[i] = angka[i+1]
	}

	n--

	for i := 0; i < n; i++ {
		fmt.Println(angka[i])
	}
}
