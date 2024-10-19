package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= b {
		cetakPermutasi(a, b)
	} else {
		cetakPermutasi(b, a)
	}
}

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Prosedur untuk mencetak permutasi
func cetakPermutasi(n, r int) {
	perm := permutasi(n, r)
	fmt.Println(perm)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}
