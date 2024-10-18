// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Fungsi utama program
func main() {
	var a, b int
	// Menerima input dari pengguna
	fmt.Scan(&a, &b)
	// Memeriksa apakah a >= b
	if a >= b {
		// Jika ya, panggil prosedur hitungPermutasi(a, b)
		hitungPermutasi(a, b)
	} else {
		// Jika tidak, panggil prosedur hitungPermutasi(b, a)
		hitungPermutasi(b, a)
	}
}

// Prosedur untuk menghitung faktorial dari n
func hitungFaktorial(n int) {
	var hasil int = 1
	var i int
	// Loop untuk menghitung faktorial
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	fmt.Println("Faktorial dari", n, "adalah", hasil)
}

// Prosedur untuk menghitung permutasi P(n, r)
func hitungPermutasi(n, r int) {
	hitungFaktorial(n)
	hitungFaktorial(n - r)
	var hasil int = 1
	var i int
	// Loop untuk menghitung faktorial n
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	var hasil2 int = 1
	// Loop untuk menghitung faktorial n-r
	for i = 1; i <= n-r; i++ {
		hasil2 = hasil2 * i
	}
	fmt.Println("Permutasi dari", n, "dan", r, "adalah", hasil/hasil2)
}
