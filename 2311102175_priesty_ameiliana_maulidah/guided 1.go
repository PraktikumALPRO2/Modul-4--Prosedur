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

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(n, r int) {
	result := faktorial(n) / faktorial(n-r)
	fmt.Println(result)
}
