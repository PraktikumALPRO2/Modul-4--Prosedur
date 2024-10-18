
package main

import (
	"fmt"
)

// Prosedur untuk menghitung dan menampilkan luas persegi
func hitungLuas(sisi float64) {
	luas := sisi * sisi
	fmt.Printf("Luas persegi: %.2f\n", luas)
}

// Prosedur untuk menghitung dan menampilkan keliling persegi
func hitungKeliling(sisi float64) {
	keliling := 4 * sisi
	fmt.Printf("Keliling persegi: %.2f\n", keliling)
}

func main() {
	var sisi float64

	fmt.Print("Sisi Persegi: ")
	fmt.Scan(&sisi)

	// Panggil prosedur untuk menghitung luas dan keliling
	hitungLuas(sisi)
	hitungKeliling(sisi)
}
