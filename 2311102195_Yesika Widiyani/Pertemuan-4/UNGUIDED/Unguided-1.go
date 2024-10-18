package main

import "fmt"

// Prosedur untuk membaca input dari pengguna
func bacaInput(a *int, b *int) {
	fmt.Scan(a, b)
}

// Prosedur untuk mencetak hasil permutasi
func cetakPermutasi(n, r int) {
	if n >= r {
		fmt.Println(permutasi(n, r))
	} else {
		fmt.Println(permutasi(r, n))
	}
}

// Fungsi faktorial untuk menghitung nilai faktorial dari suatu bilangan
func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi permutasi untuk menghitung permutasi nPr
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func main() {
	var a, b int
	// Prosedur untuk membaca input
	bacaInput(&a, &b)
	// Prosedur untuk mencetak hasil permutasi
	cetakPermutasi(a, b)
}
