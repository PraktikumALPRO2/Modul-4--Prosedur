package main

import "fmt"

// Prosedur untuk menghitung dan mencetak luas persegi
func hitungLuas(sisi float64) {
	luas := sisi * sisi
	fmt.Printf("Luas persegi: %.2f\n", luas) // Mencetak hasil luas
}

// Prosedur untuk menghitung dan mencetak keliling persegi
func hitungKeliling(sisi float64) {
	keliling := 4 * sisi
	fmt.Printf("Keliling persegi: %.2f\n", keliling) // Mencetak hasil keliling
}

func main() {
	var sisi float64

	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	hitungLuas(sisi)     // Memanggil prosedur hitungLuas
	hitungKeliling(sisi) // Memanggil prosedur hitungKeliling
}
