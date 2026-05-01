package main

import "fmt"

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func main() {
	var angka [20]int
	var n int

	fmt.Scan(&n)

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan angka %d: ", i)
		fmt.Scan(&angka[i])
	}

	total := 0
	for i := 0; i < n; i++ {
		total += angka[i]
	}
	rata := float64(total) / float64(n)

	var jumlah float64
	for i := 0; i < n; i++ {
		selisih := float64(angka[i]) - rata
		jumlah += selisih * selisih
	}

	stdDev := sqrt(jumlah / float64(n))

	fmt.Println(rata)
	fmt.Println(stdDev)
}
