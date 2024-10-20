package main

import "fmt"

// Fungsi utama
func main() {
	var a, b int
	fmt.Print("Masukkan dua bilangan: ")
	fmt.Scan(&a, &b)

	if a >= b {
		var hasil int
		permutasi(a, b, &hasil)
		fmt.Println("Hasil permutasi:", hasil)
	} else {
		var hasil int
		permutasi(b, a, &hasil)
		fmt.Println("Hasil permutasi:", hasil)
	}
}

// Prosedur untuk menghitung faktorial
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi
func permutasi(n, r int, hasil *int) {
	var faktN, faktNR int
	faktorial(n, &faktN)
	faktorial(n-r, &faktNR)
	*hasil = faktN / faktNR
}
