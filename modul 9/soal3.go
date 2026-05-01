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
