package main

import "fmt"

// Prosedur faktorial untuk menghitung faktorial dari n
func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur permutasi untuk menghitung permutasi P(n, r)
func permutasi(n, r int, hasil *int) {
	var faktN, faktNR int
	faktorial(n, &faktN)    // Faktorial n
	faktorial(n-r, &faktNR) // Faktorial (n-r)
	*hasil = faktN / faktNR // Permutasi P(n, r)
}

// Prosedur kombinasi untuk menghitung kombinasi C(n, r)
func kombinasi(n, r int, hasil *int) {
	var faktN, faktR, faktNR int
	faktorial(n, &faktN)              // Faktorial n
	faktorial(r, &faktR)              // Faktorial r
	faktorial(n-r, &faktNR)           // Faktorial (n-r)
	*hasil = faktN / (faktR * faktNR) // Kombinasi C(n, r)
}

func main() {
	var a, b, c, d int
	var permA, combA, permB, combB int

	// Membaca input dari pengguna
	fmt.Print("Masukkan 4 bilangan bulat (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung permutasi dan kombinasi untuk (a, c) dan (b, d)
	permutasi(a, c, &permA)
	kombinasi(a, c, &combA)
	permutasi(b, d, &permB)
	kombinasi(b, d, &combB)

	// Mencetak hasil
	fmt.Println(permA, combA) // Hasil permutasi dan kombinasi a terhadap c
	fmt.Println(permB, combB) // Hasil permutasi dan kombinasi b terhadap d
}
