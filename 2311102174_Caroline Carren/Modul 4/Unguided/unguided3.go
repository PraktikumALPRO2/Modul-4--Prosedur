// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Mencetak nilai awal
	fmt.Print(n)

	// Menghitung dan mencetak deret
	for n != 1 {
		if n%2 == 0 {
			n = n / 2 // Jika genap, bagi dua
		} else {
			n = 3*n + 1 // Jika ganjil, 3n + 1
		}
		fmt.Print(" ", n) // Mencetak suku berikutnya dengan spasi
	}
}

func main() {
	var nilaiAwal int
	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
	fmt.Scan(&nilaiAwal)

	if nilaiAwal > 0 && nilaiAwal < 1000000 {
		cetakDeret(nilaiAwal)
		fmt.Println() // Mencetak newline setelah deret
	} else {
		fmt.Println("Masukkan bilangan bulat positif yang valid.")
	}
}
