package main

import (
	"fmt"
)

// Prosedur untuk menghitung dan menampilkan luas dan keliling persegi
func hitungPersegi(sisi float64) {
	// Menghitung luas persegi
	luas := sisi * sisi

	// Menghitung keliling persegi
	keliling := 4 * sisi

	// Hasil
	fmt.Printf("Luas persegi: %.2f\n", luas)
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}

func main() {
	// Deklarasi variabel
	var sisi float64

	// Input panjang sisi persegi
	fmt.Print("Masukkan panjang sisi: ")
	fmt.Scan(&sisi)

	// Menghitung dan menampilkan luas dan keliling persegi
	hitungPersegi(sisi)
}
