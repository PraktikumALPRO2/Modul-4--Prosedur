package main

import (
	"fmt"
)

// Prosedur factorial menghitung faktorial dari sebuah angka n
func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur permutation menghitung P(n, r) = n! / (n - r)!
func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)    // Menghitung n!
	factorial(n-r, &fnr) // Menghitung (n - r)!
	*hasil = fn / fnr    // Hasil permutasi
}

// Prosedur combination menghitung C(n, r) = n! / (r! * (n - r)!)
func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)        // Menghitung n!
	factorial(r, &fr)        // Menghitung r!
	factorial(n-r, &fnr)     // Menghitung (n - r)!
	*hasil = fn / (fr * fnr) // Hasil kombinasi
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	// Memvalidasi kondisi: a >= c dan b >= d
	if a < c || b < d {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d.")
		return
	}

	// Menghitung permutasi dan kombinasi
	permutation(a, c, &p1)
	combination(a, c, &c1)
	permutation(b, d, &p2)
	combination(b, d, &c2)

	// Menampilkan hasil
	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
}
