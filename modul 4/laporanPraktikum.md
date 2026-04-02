# <h1 align="center">Laporan Praktikum Modul 4 - ... </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan_T4/blob/main/modul%204/output/output-soal1.png)
[Penjelasan: Jadi program ini digunakan untuk menghitung permutasi dan kombinasi dari dua pasangan bilangan. Program menerima input a, b, c, d, lalu menghitung P(a,c) dan C(a,c) serta P(b,d) dan C(b,d) menggunakan rumus faktorial. Hasilnya ditampilkan dalam dua baris sesuai pasangan bilangan.]

### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T4/blob/main/modul%204/output/output-soal2.png)
[Penjelasan: Jadi program ini menentukan pemenang lomba berdasarkan jumlah soal yang diselesaikan dan total waktu. Setiap peserta memiliki 8 waktu pengerjaan, di mana nilai 301 dianggap tidak selesai. Program menghitung jumlah soal selesai dan total waktu, lalu memilih pemenang dengan soal terbanyak, atau waktu tercepat jika jumlahnya sama.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for {
		fmt.Print(n)
		if n == 1 {
			break
		}

		fmt.Print(" ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T4/blob/main/modul%204/output/output-soal3.png)
[Penjelasan: Jadi program ini menghasilkan deret bilangan dari suatu angka hingga mencapai 1. Jika bilangan genap dibagi 2, jika ganjil dikali 3 lalu ditambah 1. Proses diulang sampai 1 dan semua nilai dicetak dalam satu baris.]



