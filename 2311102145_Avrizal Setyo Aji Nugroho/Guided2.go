package main

import (
	"fmt"
)

// Prosedur untuk menghitung luas
func Luas(s int, L *int) {
	*L = s * s
}

// Prosedur untuk menghitung keliling
func Keliling(s int, K *int) {
	*K = 4 * s
}

func main() {
	var s int
	var L, K int

	fmt.Print("Inputkan sisi: ")
	fmt.Scan(&s)

	// Memanggil prosedur untuk menghitung luas dan keliling
	Luas(s, &L)
	Keliling(s, &K)

	// Menampilkan hasil luas dan keliling
	fmt.Println("Luas: ", L)
	fmt.Println("Keliling: ", K)
}
