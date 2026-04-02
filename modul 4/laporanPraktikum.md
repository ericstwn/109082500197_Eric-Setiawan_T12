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
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan/blob/main/modul1/output/output-soal1.png)
[Penjelasan: Jadi program ini berguna untuk merubah posisi dari variabelnya dimana temp
menyimpan nomor satu,satu menyimpan variabel nomor dua,dua menyimpan variabel
nomor tiga,dan tiga menyimpan variabel punya temp dimana tempnya sudah diganti dengan variabel nomor satu.]

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
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan/blob/main/modul1/output/output-soal2.png)
[Penjelasan: Jadi Program ini memakai perulangan untuk input empat warna sebanyak lima kali,
lalu membandingkan input tersebut dengan urutan standar atau true "merah", "kuning",
"hijau", dan "ungu". Variabel berhasil itu sebagai true, namun akan diubah
menjadi false jika ditemukan satu saja percobaan yang urutan warnanya tidak sesuai, sehingga
program akan menghasilkan nilai true jika lima percobaan memiliki
urutan yang tepat.]

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
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan/blob/main/modul1/output/output-soal3.png)
[Penjelasan: Jadi program ini menghitung biaya pengiriman parsel dengan membagi total berat gram
menjadi jumlah kilogram dan sisa gram, di mana biaya dasar ditetapkan sebesar Rp10.000
per kilogram. Program kemudian menerapkan aturan khusus: jika berat total kurang dari 10
kg, sisa gram akan dikenakan biaya Rp5 per gram jika jumlahnya 500 gram atau lebih, dan
Rp15 per gram jika kurang dari 500 gram; namun, jika berat total sudah melebihi 10 kg,
maka sisa gram tersebut digratiskan, sehingga total biaya hanya dihitung berdasarkan berat
kilogramnya saja.]



