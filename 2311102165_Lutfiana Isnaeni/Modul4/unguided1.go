// Lutfiana Isnaeni Lathifah
// 2311102165
package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n - r)!
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r! * (n - r)!)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	// Input empat buah bilangan a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan empat bilangan (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Asumsikan syarat a >= c dan b >= d sudah terpenuhi

	// Hitung permutasi dan kombinasi untuk (a, c) dan (b, d)
	permA := permutasi(a, c)
	kombA := kombinasi(a, c)
	permB := permutasi(b, d)
	kombB := kombinasi(b, d)

	// Cetak hasil
	fmt.Printf("Permutasi(%d, %d): %d\n", a, c, permA)
	fmt.Printf("Kombinasi(%d, %d): %d\n", a, c, kombA)
	fmt.Printf("Permutasi(%d, %d): %d\n", b, d, permB)
	fmt.Printf("Kombinasi(%d, %d): %d\n", b, d, kombB)
}
