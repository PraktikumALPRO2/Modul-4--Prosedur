package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		permutasi(a, b)
	} else {
		permutasi(b, a)
	}
}

// Prosedur faktorial
func faktorial(n int, result *int) {
	*result = 1
	for i := 1; i <= n; i++ {
		*result *= i
	}
}

// Prosedur permutasi
func permutasi(n, r int) {
	var faktN, faktNR int
	faktorial(n, &faktN)      // Hitung faktorial(n)
	faktorial(n-r, &faktNR)   // Hitung faktorial(n-r)
	result := faktN / faktNR  // Hasil perhitungan permutasi
	fmt.Println(result)       // Cetak hasil
}
