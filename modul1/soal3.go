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
