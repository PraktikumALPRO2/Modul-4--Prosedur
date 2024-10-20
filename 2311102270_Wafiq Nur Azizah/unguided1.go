package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		var permAC, combAC, permBD, combBD int

		permutation(a, c, &permAC)
		combination(a, c, &combAC)

		permutation(b, d, &permBD)
		combination(b, d, &combBD)

		fmt.Println(permAC, combAC)
		fmt.Println(permBD, combBD)
	} else {
		fmt.Println("Syarat tidak terpenuhi.")
	}
}