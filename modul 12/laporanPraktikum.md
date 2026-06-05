# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func hitungSuara(T tabInt, n int, calon int) int {
	count := 0
	for i := 0; i < n; i++ {
		if T[i] == calon {
			count++
		}
	}
	return count
}

func main() {
	var T tabInt
	var x int
	var n int
	var suaraMasuk int
	var suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			T[n] = x
			n++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println()

	for calon := 1; calon <= 20; calon++ {
		jumlah := hitungSuara(T, n, calon)
		if jumlah > 0 {
			fmt.Printf("%d: %d\n", calon, jumlah)
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan_T10/blob/main/modul%2010/output/output-soal1.png)
[Penjelasan: Jadi program ini buat ngitung hasil pemilihan suara dari 20 calon. Program terus minta input nomor calon sampai diinput 0 sebagai tanda selesai. Setiap suara yang masuk dihitung, tapi hanya suara dengan nomor 1 sampai 20 yang dianggap sah, selain itu dianggap tidak sah. Fungsi hitungSuara ngitung berapa banyak suara yang diterima tiap calon. Hasilnya ditampilin total suara masuk, total suara sah, lalu jumlah suara tiap calon yang mendapat suara.]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func hitungSuara(T tabInt, n int, calon int) int {
	count := 0

	for i := 0; i < n; i++ {
		if T[i] == calon {
			count++
		}
	}

	return count
}

func main() {
	var T tabInt
	var x int
	var n int
	var suaraMasuk int
	var suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			T[n] = x
			n++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	ketua := 1
	wakil := 2

	for calon := 1; calon <= 20; calon++ {

		if hitungSuara(T, n, calon) > hitungSuara(T, n, ketua) {
			wakil = ketua
			ketua = calon

		} else if calon != ketua &&
			(hitungSuara(T, n, calon) > hitungSuara(T, n, wakil) ||
				wakil == ketua) {

			wakil = calon
		}
	}

	if hitungSuara(T, n, wakil) > hitungSuara(T, n, ketua) {
		ketua, wakil = wakil, ketua
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T10/blob/main/modul%2010/output/output-soal2.png)
[Penjelasan: Jadi program ini buat ngitung hasil pemilihan ketua dan wakil ketua RT. Program terus minta input nomor calon sampai diinput 0 sebagai tanda selesai. Suara yang sah hanya nomor 1 sampai 20, selain itu tidak dihitung. Setelah semua suara masuk, program nyari siapa yang mendapat suara terbanyak sebagai ketua, dan suara terbanyak kedua sebagai wakil ketua, dengan cara bandingin perolehan suara semua calon satu-satu. Hasilnya ditampilin total suara masuk, total suara sah, nomor calon ketua RT, dan nomor calon wakil ketua RT.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	return -1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T10/blob/main/modul%2010/output/output-soal3.png)
[Penjelasan: Jadi program ini buat nyari posisi (indeks) suatu angka di dalam kumpulan data yang sudah berurutan. Di awal, program bakal minta input jumlah total data dan angka yang mau dicari, baru setelah itu kita masukin semua isi datanya satu per satu. Proses nyarinya pakai metode Binary Search, yaitu dengan cara langsung ngecek angka di posisi paling tengah; kalau angka yang dicari kekecilan atau kegedean, setengah datanya langsung dibuang dan sisanya dibelah dua lagi sampai ketemu. Hasil akhirnya bakal nampilin nomor posisi angka tersebut, atau tulisan "TIDAK ADA" kalau angkanya emang gak ketemu di dalam data.]





