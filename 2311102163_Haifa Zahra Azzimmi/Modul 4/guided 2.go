package main

import (
	"fmt"
)

// Prosedur utama
func main() {
	var panjangSisi float64

	// Meminta input panjang sisi persegi
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Menghitung luas dan keliling persegi menggunakan prosedur terpisah
	luas := hitungLuas(panjangSisi)
	keliling := hitungKeliling(panjangSisi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", luas)
	fmt.Printf("Keliling persegi: %g\n", keliling)
}

// Prosedur untuk menghitung luas persegi
func hitungLuas(sisi float64) float64 {
	return sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKeliling(sisi float64) float64 {
	return 4 * sisi
}
