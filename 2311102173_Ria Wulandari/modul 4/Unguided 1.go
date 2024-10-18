package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi P(n, r) = n! / (n-r)!
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r) = n! / (r!(n-r)!)
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// Input
	fmt.Println("Masukkan empat bilangan asli (a, b, c, d):")
	fmt.Scan(&a, &b, &c, &d)

	// Baris pertama: permutasi dan kombinasi a terhadap c
	perm1 := permutation(a, c)
	comb1 := combination(a, c)

	// Baris kedua: permutasi dan kombinasi b terhadap d
	perm2 := permutation(b, d)
	comb2 := combination(b, d)

	// Output
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, perm1, comb1)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, perm2, comb2)
}
