package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial dari sebuah bilangan
func hitungFaktorial(n int) int {
	if n == 0 {
		return 1
	}
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r)
func hitungPermutasi(n, r int) int {
	faktorialN := hitungFaktorial(n)
	faktorialNR := hitungFaktorial(n - r)
	return faktorialN / faktorialNR
}

// Fungsi untuk menghitung kombinasi C(n, r)
func hitungKombinasi(n, r int) int {
	faktorialN := hitungFaktorial(n)
	faktorialR := hitungFaktorial(r)
	faktorialNR := hitungFaktorial(n - r)
	return faktorialN / (faktorialR * faktorialNR)
}

func main() {
	// Input
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Tampilkan hasil permutasi dan kombinasi untuk a dan c
	fmt.Printf("Permutasi P(%d, %d) adalah: %d\n", a, c, hitungPermutasi(a, c))
	fmt.Printf("Kombinasi C(%d, %d) adalah: %d\n", a, c, hitungKombinasi(a, c))

	// Tampilkan hasil permutasi dan kombinasi untuk b dan d
	fmt.Printf("Permutasi P(%d, %d) adalah: %d\n", b, d, hitungPermutasi(b, d))
	fmt.Printf("Kombinasi C(%d, %d) adalah: %d\n", b, d, hitungKombinasi(b, d))
}
