package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a >= b {
		var result int
		permutasi(a, b, &result)
		fmt.Print(result)
	} else {
		var result int
		permutasi(b, a, &result)
		fmt.Print(result)
	}
}

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, result *int) {
	var faktN, faktNR int
	faktorial(n, &faktN)
	faktorial(n-r, &faktNR)
	*result = faktN / faktNR
}
