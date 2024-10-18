package main

import (
	"fmt"
)

// Fungsi untuk mencetak deret berdasarkan nilai n
func cetakDeret(n int) {
	// Teruskan pencetakan deret selama n tidak sama dengan 1
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n /= 2 // Jika n adalah genap
		} else {
			n = 3*n + 1 // Jika n adalah ganjil
		}
	}
	// Cetak 1 sebagai elemen terakhir dari deret
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Memanggil fungsi untuk mencetak deret
	cetakDeret(n)
}
