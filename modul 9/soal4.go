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
