package main

import "fmt"

// Prosedur untuk membaca input dari pengguna
func bacaInput(a, b *int) {
	fmt.Scan(a, b)
}

// Prosedur untuk menghitung faktorial
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung dan mencetak permutasi
func hitungPermutasi(n, r int) {
	var faktorialN, faktorialNR int

	// Menghitung faktorial n
	faktorial(n, &faktorialN)

	// Menghitung faktorial (n-r)
	faktorial(n-r, &faktorialNR)

	// Menghitung permutasi
	hasil := faktorialN / faktorialNR
	fmt.Println(hasil)
}

// Prosedur utama yang memutuskan dan memanggil hitungPermutasi
func permutasi(a, b int) {
	if a >= b {
		hitungPermutasi(a, b)
	} else {
		hitungPermutasi(b, a)
	}
}

func main() {
	var a, b int

	// Memanggil prosedur bacaInput untuk mendapatkan nilai a dan b
	bacaInput(&a, &b)

	// Memanggil prosedur permutasi
	permutasi(a, b)
}
