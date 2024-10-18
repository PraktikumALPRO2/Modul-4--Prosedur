package main

import (
	"fmt"
)

// Prosedur untuk menghitung faktorial
func HitungFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung permutasi P(n, r)
func HitungPermutasi(n, r int, hasil *int) {
	var fn, fnr int
	HitungFaktorial(n, &fn)
	HitungFaktorial(n-r, &fnr)
	*hasil = fn / fnr
}

// Prosedur untuk menghitung kombinasi C(n, r)
func HitungKombinasi(n, r int, hasil *int) {
	var fn, fnr, fr int
	HitungFaktorial(n, &fn)
	HitungFaktorial(n-r, &fnr)
	HitungFaktorial(r, &fr)
	*hasil = fn / (fnr * fr)
}

func main() {
	// Input empat bilangan asli a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan empat bilangan asli (a, b, c, d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Syarat: a >= c dan b >= d
	if a < c || b < d {
		fmt.Println("Syarat tidak terpenuhi: a harus >= c dan b harus >= d.")
		return
	}

	// Variabel untuk hasil permutasi dan kombinasi
	var hasilPermutasiA, hasilKombinasiA, hasilPermutasiB, hasilKombinasiB int

	// Hitung permutasi dan kombinasi untuk a terhadap c
	HitungPermutasi(a, c, &hasilPermutasiA)
	HitungKombinasi(a, c, &hasilKombinasiA)

	// Hitung permutasi dan kombinasi untuk b terhadap d
	HitungPermutasi(b, d, &hasilPermutasiB)
	HitungKombinasi(b, d, &hasilKombinasiB)

	// Output hasil
	fmt.Printf("\nPermutasi dan Kombinasi untuk %d dan %d: %d %d\n", a, c, hasilPermutasiA, hasilKombinasiA)
	fmt.Printf("Permutasi dan Kombinasi untuk %d dan %d: %d %d\n", b, d, hasilPermutasiB, hasilKombinasiB)
}
