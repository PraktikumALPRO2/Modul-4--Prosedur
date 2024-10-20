package main

import "fmt"

func Faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func Permutasi(n, r int, hasil *int) {
	var fn, fnr int
	Faktorial(n, &fn)
	Faktorial(n-r, &fnr)
	*hasil = fn / fnr
}

func Kombinasi(n, r int, hasil *int) {
	var fn, fr, fnr int
	Faktorial(n, &fn)
	Faktorial(r, &fr)
	Faktorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a*c >= d && b >= d {
		var permAC, combAC, permBD, combBD int
		Permutasi(a, c, &permAC)
		Kombinasi(a, c, &combAC)
		Permutasi(b, d, &permBD)
		Kombinasi(b, d, &combBD)

		fmt.Println(permAC, combAC)
		fmt.Println(permBD, combBD)
	} else {
		fmt.Println("Syarat tidak terpenuhi.")
	}
}
