# <h1 align="center">Laporan Praktikum Modul 5 - ... </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(fibo(i), " ")
	}
}

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal1.png)
[Penjelasan: Jadi program ini digunakan untuk mencetak deret bilangan Fibonacci sebanyak n suku menggunakan rekursi. Program menerima input n, lalu memanggil fungsi fibo(i) untuk setiap indeks dari 0 hingga n-1. Fungsi fibo bekerja secara rekursif: jika n = 0 mengembalikan 0, jika n = 1 mengembalikan 1, dan untuk nilai lainnya mengembalikan hasil penjumlahan fibo(n-1) + fibo(n-2).]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}

func bintang(n int) {
	if n == 0 {
		return
	}
	bintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal2.png)
[Penjelasan: Jadi program ini digunakan untuk mencetak pola segitiga bintang secara rekursif. Program menerima input n, lalu fungsi bintang dipanggil secara rekursif hingga n = 0. Karena pemanggilan rekursif dilakukan sebelum mencetak bintang, baris pertama akan mencetak 1 bintang, baris kedua 2 bintang, dan seterusnya hingga baris ke-n yang mencetak n bintang, membentuk pola segitiga yang semakin besar.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
func faktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	faktor(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal3.png)
[Penjelasan: Jadi program ini digunakan untuk mencari dan mencetak semua faktor dari bilangan n menggunakan rekursi. Program menerima input n, lalu fungsi faktor dipanggil mulai dari i = 1. Setiap kali i habis membagi n (n % i == 0), nilai i dicetak sebagai faktor. Proses berlanjut secara rekursif dengan menambah i sebesar 1 hingga i melebihi n.]

### 4. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	test(n)
}

func test(n int) {
	if n == 0 {
		return
	} else {
		fmt.Print(n)
		test(n - 1)
		fmt.Print(n)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal4.png)
[Penjelasan: Jadi program ini digunakan untuk mencetak angka secara maju lalu mundur menggunakan rekursi. Program menerima input n, lalu fungsi test mencetak n sebelum rekursi dan kembali mencetak n setelah rekursi selesai. Hasilnya adalah barisan angka dari n turun ke 1, kemudian naik kembali dari 1 ke n, misalnya jika n = 3 maka outputnya adalah 3 2 1 1 2 3.]

### 5. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bilangan(1, n)
}

func bilangan(i int, n int) {
	if i > n {
		return
	}
	fmt.Print(i, " ")
	bilangan(i+2, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal5.png)
[Penjelasan: Jadi program ini digunakan untuk mencetak bilangan ganjil dari 1 hingga n menggunakan rekursi. Program menerima input n, lalu fungsi bilangan dipanggil mulai dari i = 1 dan setiap rekursi menambahkan i sebesar 2. Dengan demikian hanya bilangan ganjil (1, 3, 5, 7, ...) yang akan dicetak selama i masih kurang dari atau sama dengan n.]

### 6. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	pangkat(x, y)
	fmt.Println(pangkat(x, y))
}

func pangkat(x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pangkat(x, y-1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T5/blob/main/modul%205/output/output-soal6.png)
[Penjelasan: Jadi program ini digunakan untuk menghitung hasil perpangkatan x pangkat y menggunakan rekursi. Program menerima input x dan y, lalu fungsi pangkat bekerja secara rekursif: jika y = 0 mengembalikan 1 (karena x^0 = 1), dan selain itu mengembalikan x dikali pangkat(x, y-1).]



