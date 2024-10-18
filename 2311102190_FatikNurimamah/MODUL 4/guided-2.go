package main

import (
	"fmt"
)

func main() {
	// Meminta input panjang sisi persegi
	var panjangSisi float64
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&panjangSisi)

	var luasPersegi, kelilingPersegi float64
	// Memanggil prosedur untuk menghitung luas dan keliling persegi
	hitungLuasPersegi(panjangSisi, &luasPersegi)
	hitungKelilingPersegi(panjangSisi, &kelilingPersegi)

	// Menampilkan hasil
	fmt.Printf("Luas persegi: %g\n", luasPersegi)
	fmt.Printf("Keliling persegi: %g\n", kelilingPersegi)
}

// Prosedur untuk menghitung luas persegi
func hitungLuasPersegi(sisi float64, hasil *float64) {
	*hasil = sisi * sisi
}

// Prosedur untuk menghitung keliling persegi
func hitungKelilingPersegi(sisi float64, hasil *float64) {
	*hasil = 4 * sisi
}
