package main

import "fmt"

func main() {
	// Mendeklarasikan variabel untuk menampung dua input bilangan bulat
	var a, b int
	fmt.Scan(&a, &b)

	// Mendeklarasikan variabel untuk menampung hasil permutasi
	var hasilPermutasi int
	// Memanggil prosedur hitungPermutasi berdasarkan nilai a dan b
	if a >= b {
		hitungPermutasi(a, b, &hasilPermutasi)
	} else {
		hitungPermutasi(b, a, &hasilPermutasi)
	}

	// Menampilkan hasil permutasi
	fmt.Println(hasilPermutasi)
}

// Prosedur untuk menghitung faktorial dari bilangan n
func hitungFaktorial(n int, hasil *int) {
	*hasil = 1
	// Melakukan perulangan untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi dari n dan r
func hitungPermutasi(n, r int, hasil *int) {
	var fn, fnr int
	hitungFaktorial(n, &fn)
	hitungFaktorial(n-r, &fnr)
	// Menghitung permutasi dengan rumus n! / (n-r)! dan menyimpannya di hasil
	*hasil = fn / fnr
}
