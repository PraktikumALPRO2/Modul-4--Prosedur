package main

import "fmt"

// Prosedur utama
func main() {
	var a, b int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	// Memilih mana yang lebih besar antara a dan b untuk menghitung permutasi
	if a >= b {
		fmt.Printf("Permutasi dari %dP%d adalah: %d\n", a, b, permutasi(a, b))
	} else {
		fmt.Printf("Permutasi dari %dP%d adalah: %d\n", b, a, permutasi(b, a))
	}
}

// Prosedur untuk menghitung faktorial dari suatu bilangan
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Prosedur untuk menghitung permutasi nPr
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
