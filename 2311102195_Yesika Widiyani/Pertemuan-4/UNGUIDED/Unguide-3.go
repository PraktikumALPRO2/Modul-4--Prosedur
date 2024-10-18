package main

import "fmt"

// Prosedur untuk mencetak deret berdasarkan nilai awal
func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ") // Mencetak nilai n diikuti spasi
		if n%2 == 0 {
			n = n / 2 // Jika n genap, bagi 2
		} else {
			n = 3*n + 1 // Jika n ganjil, 3n + 1
		}
	}
	fmt.Print(n) // Mencetak 1 sebagai suku terakhir
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif (kurang dari 1000000): ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n >= 1000000 {
		fmt.Println("Input tidak valid! Masukkan bilangan bulat positif kurang dari 1000000.")
		return
	}

	// Mencetak deret
	cetakDeret(n)
}
