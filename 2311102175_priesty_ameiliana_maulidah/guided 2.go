package main

import (
	"fmt"
)

func hitungPersegi(sisi float64) (float64, float64) {
	luas := sisi * sisi
	keliling := 4 * sisi
	return luas, keliling
}

func main() {
	var sisi float64
	fmt.Print("masukkan panjang sisi persegi: ")
	fmt.Scan(&sisi)

	luas, keliling := hitungPersegi(sisi)

	fmt.Printf("luas persegi: %.2f\n", luas)
	fmt.Printf("keliling persegi: %.2f\n", keliling)
}
