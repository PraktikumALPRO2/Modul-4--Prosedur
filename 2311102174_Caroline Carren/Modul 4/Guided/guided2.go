// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import "fmt"

// Prosedur untuk menghitung luas persegi
func hitungLuas(sisi float64) {
	var luas float64 = sisi * sisi
	fmt.Printf("Luas persegi : %.2f\n", luas)
}

// Prosedur untuk menghitung keliling persegi
func hitungKeliling(sisi float64) {
	var keliling float64 = 4 * sisi
	fmt.Printf("Keliling persegi : %.2f\n", keliling)
}

func main() {
	var sisi float64

	// Tanya pengguna untuk memasukkan sisi persegi
	fmt.Print("Masukkan sisi persegi : ")
	fmt.Scanln(&sisi)

	// Panggil prosedur untuk menghitung luas dan keliling
	hitungLuas(sisi)
	hitungKeliling(sisi)
}
