package main

import (
	"fmt"
)

// Fungsi untuk menghitung faktorial
func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutasi(n int, r int, hasil *int) {
	var fn, fn_r int
	factorial(n, &fn)
	factorial(n-r, &fn_r)
	*hasil = fn / fn_r
}

// Fungsi untuk menghitung kombinasi C(n, r)
func kombinasi(n int, r int, hasil *int) {
	var fn, fr, fn_r int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fn_r)
	*hasil = fn / (fr * fn_r)
}

func main() {
	var a, b, c, d int

	// Meminta input dari pengguna dalam satu baris
	fmt.Print("Masukkan nilai a, b, c, dan d : ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Cek syarat a ≥ c dan b ≥ d
	if a >= c && b >= d {
		var hasilPermutasi1, hasilKombinasi1, hasilPermutasi2, hasilKombinasi2 int

		// Hitung permutasi dan kombinasi untuk a, c
		permutasi(a, c, &hasilPermutasi1)
		kombinasi(a, c, &hasilKombinasi1)

		// Hitung permutasi dan kombinasi untuk b, d
		permutasi(b, d, &hasilPermutasi2)
		kombinasi(b, d, &hasilKombinasi2)

		// Tampilkan hasil
		fmt.Println("Hasil:")
		fmt.Printf("P(%d,%d) = %d\n", a, c, hasilPermutasi1)
		fmt.Printf("C(%d,%d) = %d\n", a, c, hasilKombinasi1)
		fmt.Printf("P(%d,%d) = %d\n", b, d, hasilPermutasi2)
		fmt.Printf("C(%d,%d) = %d\n", b, d, hasilKombinasi2)
	} else {
		fmt.Println("Syarat a ≥ c dan b ≥ d tidak terpenuhi")
	}
}
