package main

import "fmt"

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

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var hasilPermutasi int

	if a >= b {
		permutasi(a, b, &hasilPermutasi)
	} else {
		permutasi(b, a, &hasilPermutasi)
	}

	fmt.Println(hasilPermutasi)
}
