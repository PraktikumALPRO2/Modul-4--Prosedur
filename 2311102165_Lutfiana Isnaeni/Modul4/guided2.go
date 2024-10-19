package main

import (
	"fmt"
)

// Prosedur untuk menghitung luas dan keliling persegi
func hitungPersegi(panjangSisi float64) (float64, float64) {
	// Menghitung luas dan keliling persegi
	LuasPersegi := panjangSisi * panjangSisi
	KelilingPersegi := 4 * panjangSisi

	return LuasPersegi, KelilingPersegi
}

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	// Memanggil prosedur untuk menghitung luas dan keliling
	Luas, Keliling := hitungPersegi(panjangSisi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", Luas)
	fmt.Printf("Keliling persegi: %g\n", Keliling)
}
