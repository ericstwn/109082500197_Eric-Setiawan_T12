package main

import "fmt"

func hitungSkor(soal *int, skor *int, waktu []int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		if waktu[i] <= 300 {
			*soal++
			*skor += waktu[i]
		}
	}
}

func main() {
	var nama string

	pemenang := ""
	maxSoal := -1
	minWaktu := 1<<31 - 1

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Scan(&waktu[i])
		}

		var soal, skor int
		hitungSkor(&soal, &skor, waktu)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}
	fmt.Println(pemenang, maxSoal, minWaktu)
}
