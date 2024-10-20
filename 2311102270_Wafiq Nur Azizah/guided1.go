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

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ { // Perbaikan di sini: `i <= n`
		*hasil *= i
	}
}

func permutasi(n, r int) {
	var hasiln, hasilnr int
	faktorial(n, &hasiln)
	faktorial(n-r, &hasilnr)
	fmt.Println("Hasil permutasi:", hasiln/hasilnr)
}
