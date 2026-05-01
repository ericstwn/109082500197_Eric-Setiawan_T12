# <h1 align="center">Laporan Praktikum Modul 9 - ... </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r int
}

func dalamLingkaran(t Titik, l Lingkaran) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y
	return dx*dx+dy*dy <= l.r*l.r
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	dalam1 := dalamLingkaran(t, l1)
	dalam2 := dalamLingkaran(t, l2)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal1.png)
[Penjelasan: Jadi program ini buat ngecek apakah sebuah titik berada di dalam lingkaran atau tidak, mirip program sebelumnya tapi kali ini pakai struct yang lebih terstruktur. Ada struct Titik buat nyimpen koordinat x dan y, dan struct Lingkaran yang isinya titik pusat dan jari-jari. Fungsi dalamLingkaran ngecek posisi titik dengan ngitung jarak kuadrat ke pusat lingkaran, kalau hasilnya kurang dari atau sama dengan r² berarti titiknya ada di dalam. Program ngecek dua lingkaran sekaligus dan nampilin hasilnya, titik bisa ada di dalam keduanya, salah satu, atau di luar keduanya.]

### 2. [Soal]
#### soal2 a.go

```go
package main

import "fmt"

func main() {
	var angka [5]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	angka[3] = 40
	angka[4] = 50

	fmt.Println(angka)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20a.png)
[Penjelasan: Program ini buat nyimpen dan nampilin sekumpulan angka menggunakan array. Array angka punya 5 slot yang diisi satu-satu mulai dari indeks 0 sampai 4 dengan nilai 10, 20, 30, 40, dan 50. Setelah semua slot terisi, langsung ditampilin semuanya sekaligus.]

### 2. [Soal]
#### soal2 b.go

```go
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

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Println(angka[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20b.png)
[Penjelasan: Jadi program ini buat nampilin angka-angka yang ada di posisi indeks ganjil dari array. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah semua terisi, program ngecek setiap indeks, kalau indeksnya ganjil (1, 3, 5, ...) maka angka di posisi itu ditampilin.]

### 2. [Soal]
#### soal2 c.go

```go
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

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println(angka[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20c.png)
[Penjelasan: Jadi program ini buat nampilin angka-angka yang ada di posisi indeks genap dari array. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah semua terisi, program ngecek setiap indeks, kalau indeksnya genap (0, 2, 4, ...) maka angka di posisi itu ditampilin.]

### 2. [Soal]
#### soal2 d.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20d.png)
[Penjelasan: Jadi Program ini buat nampilin angka-angka yang ada di posisi indeks kelipatan x dari array. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah itu minta input nilai x, kalau x bernilai 0 program langsung berhenti biar tidak error pembagian. Kalau x valid, program ngecek setiap indeks dan nampilin angka yang indeksnya merupakan kelipatan x (0, x, 2x, 3x, ...).]

### 2. [Soal]
#### soal2 e.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20e.png)
[Penjelasan: Jadi  Program ini buat menghapus sebuah elemen dari array berdasarkan indeks yang dipilih. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah itu minta input indeks yang mau dihapus, lalu semua elemen setelah indeks tersebut digeser satu posisi ke kiri untuk menutup celah bekas elemen yang dihapus. Jumlah elemennya dikurangi satu, lalu array yang sudah diperbarui ditampilin.]

### 2. [Soal]
#### soal2 f.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20f.png)
[Penjelasan: Jadi  Program ini buat ngitung rata-rata dari sekumpulan angka yang disimpan di array. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah semua terisi, semua angka dijumlahin ke variabel total, lalu dibagi n buat dapet nilai rata-ratanya.]

### 2. [Soal]
#### soal2 g.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20g.png)
[Penjelasan: Jadi Program ini buat ngitung rata-rata dan standar deviasi dari sekumpulan angka. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Pertama dihitung rata-ratanya dulu dengan cara menjumlahkan semua angka lalu dibagi n. Setelah itu dihitung standar deviasinya dengan cara ngitung selisih setiap angka dengan rata-rata, dikuadratkan, dijumlahin semua, dibagi n, lalu diakarkan pakai fungsi sqrt sendiri. Standar deviasi ini buat ngukur seberapa jauh angka-angka tersebut menyebar dari rata-ratanya.]

### 2. [Soal]
#### soal2 h.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20h.png)
[Penjelasan: Jadi Program ini buat ngitung berapa kali sebuah angka muncul di dalam array. Program minta input n sebagai jumlah angka, lalu minta input angkanya satu-satu dan disimpan ke array. Setelah itu minta input angka yang mau dicari, lalu program ngecek satu-satu setiap elemen array, kalau ketemu angka yang sama maka frekuensi ditambah 1. Hasilnya ditampilin berapa kali angka tersebut muncul.]


### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	var pemenang [100]string
	index := 0

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	i := 1

	fmt.Printf("Pertandingan %d : ", i)
	fmt.Scan(&skorA, &skorB)

	for skorA >= 0 && skorB >= 0 {
		if skorA > skorB {
			pemenang[index] = klubA
		} else if skorB > skorA {
			pemenang[index] = klubB
		} else {
			pemenang[index] = "Draw"
		}

		index++
		i++

		fmt.Printf("pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)
	}

	for j := 0; j < index; j++ {
		fmt.Printf("hasil %d : %s\n", j+1, pemenang[j])
	}

	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal3.png)
[Penjelasan: Jadi Program ini buat nyimpen dan nampilin hasil pertandingan antara dua klub secara berulang. Program minta input nama klub A dan klub B, lalu terus minta input skor tiap pertandingan sampai salah satu skor negatif sebagai tanda berhenti. Setiap pertandingan dicek siapa yang menang, kalau skor A lebih besar maka klub A menang, kalau skor B lebih besar maka klub B menang, kalau sama maka hasilnya Draw, dan hasilnya disimpan ke array pemenang. Setelah selesai, semua hasil pertandingan ditampilin satu-satu.]

### 4. [Soal]
#### soal4.go

```go
package main

import "fmt"

func isiArray(t *[127]rune, n *int) {
	var input string

	fmt.Print("masukkan teks: ")
	fmt.Scan(&input)

	*n = len(input)

	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
	}
}

func cetakArray(t [127]rune, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *[127]rune, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t [127]rune, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab [127]rune
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	if palindrome(tab, n) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}

	balikanArray(&tab, n)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal4.png)
[Penjelasan: Program ini buat ngecek apakah sebuah teks itu palindrom atau tidak, sekaligus nampilin teks yang sudah dibalik. Program punya empat fungsi, isiArray buat nyimpen teks ke array karakter satu-satu, cetakArray buat nampilin isi arraynya, palindrome buat ngecek apakah huruf pertama sama dengan huruf terakhir dan seterusnya sampai tengah, dan balikanArray buat nuker posisi huruf dari depan dan belakang hingga seluruh teks jadi terbalik. Jadi program ini nampilin teks asli, ngasih tau palindrom atau bukan, lalu nampilin teks yang sudah dibalik sekalian.]





