# <h1 align="center">Laporan Praktikum Modul 10 - ... </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := be
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}

		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println(min, max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal1.png)
[Penjelasan: Jadi program ini untuk mencari nilai berat terkecil dan terbesar dari sekumpulan data berat. Program minta input n sebagai jumlah data, lalu minta input beratnya satu-satu dan disimpan ke array. Nilai min dan max awalnya diset ke berat[0] sebagai pembanding pertama, lalu program ngecek satu-satu elemen berikutnya, kalau ada yang lebih kecil dari min maka min diperbarui, kalau ada yang lebih besar dari max maka max diperbarui. Hasilnya ditampilin nilai terkecil dan terbesar sekaligus.]

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func main() {
	var berat [1000]float64
	var totalWadah [1000]float64
	var x, y int

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	indeks := 0

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0

		for j := 0; j < y && indeks < x; j++ {
			total += berat[indeks]
			indeks++
		}

		totalWadah[i] = total
	}

	totalSemua := 0.0

	for i := 0; i < jumlahWadah; i++ {
		totalSemua += totalWadah[i]
	}

	rata := totalSemua / float64(jumlahWadah)

	for i := 0; i < jumlahWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(totalWadah[i])
	}

	fmt.Println()
	fmt.Println(rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20b.png)
[Penjelasan: Jadi program ini buat ngitung total berat barang per wadah dan rata-rata berat tiap wadah. Program minta input x sebagai jumlah barang dan y sebagai kapasitas tiap wadah. Semua berat barang disimpan ke array, lalu barang-barang tersebut dibagi ke dalam wadah masing-masing sebanyak y barang per wadah. Kalau sisa barangnya tidak cukup buat satu wadah penuh, tetap dibuatkan satu wadah tambahan. Total berat tiap wadah dihitung dan disimpan, lalu semua total wadah dijumlahin dan dibagi jumlah wadah buat dapet rata-ratanya. Hasilnya ditampilin total berat tiap wadah dalam satu baris lalu rata-ratanya di baris berikutnya.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int
	var min, max, rata float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &min, &max)
	rata = rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/ericstwn/109082500197_Eric-Setiawan_T9/blob/main/modul%209/output/output-soal2%20c.png)
[Penjelasan: Jadi program ini buat ngolah data berat balita, yaitu nyari berat terkecil, terbesar, dan rata-ratanya. Program minta input jumlah balita, lalu minta input berat masing-masing balita satu-satu dan disimpan ke array. Fungsi hitungMinMax nyari berat terkecil dan terbesar dengan cara bandingin satu-satu semua data, sedangkan fungsi rerata ngitung rata-rata dengan cara jumlahin semua berat lalu dibagi jumlah balitanya. Hasilnya ditampilin dengan rapi lengkap dengan satuannya yaitu kg. ]





