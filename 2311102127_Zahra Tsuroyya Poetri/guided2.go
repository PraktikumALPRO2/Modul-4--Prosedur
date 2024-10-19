package main

import "fmt"

// Fungsi untuk menghitung luas persegi
func luas(sisi int, luasPersegi *int) {
	*luasPersegi = sisi * sisi // Menghitung luas
}

// Fungsi untuk menghitung keliling persegi
func keliling(sisi int, kelPersegi *int) {
	*kelPersegi = 4 * sisi // Menghitung keliling
}

func main() {
	var sisi int
	var hasilLuas int
	var hasilKeliling int
	
	fmt.Print("Masukkan Sisi Persegi: ")
	fmt.Scan(&sisi) // Membaca input dari pengguna

	// Memanggil fungsi untuk menghitung luas dan keliling
	luas(sisi, &hasilLuas)       // Mengisi hasilLuas
	keliling(sisi, &hasilKeliling) // Mengisi hasilKeliling

	// Menampilkan hasil
	fmt.Println("Persegi dengan sisi", sisi, "memiliki luas", hasilLuas, "dan keliling", hasilKeliling)
}
