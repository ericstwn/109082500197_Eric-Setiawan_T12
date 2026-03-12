# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Eric Setiawan] - [109082500197]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
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

func main() {
	var warna1, warna2, warna3, warna4 string
	var i int
	Berhasil := true
	for i = 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d:", i)
		fmt.Scanln(&warna1, &warna2, &warna3, &warna4)
		if warna1 != "merah" || warna2 != "kuning" || warna3 != "hijau" || warna4 != "ungu" {
			Berhasil = false
		}
	}
	fmt.Println("Berhasil:", Berhasil)
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

func main() {
	var beratgram int
	fmt.Print("berat parsel: ")
	fmt.Scanln(&beratgram)

	kg := beratgram / 1000
	sisaGram := beratgram % 1000

	biayaKg := kg * 10000

	biayatambahan := 0
	if kg < 10 {
		if sisaGram >= 500 {
			biayatambahan = sisaGram * 5
		} else {
			biayatambahan = sisaGram * 15
		}
	} else {
		biayatambahan = 0
	}
	total := biayaKg + biayatambahan
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayatambahan)
	fmt.Printf("total biaya: Rp. %d\n", total)
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



