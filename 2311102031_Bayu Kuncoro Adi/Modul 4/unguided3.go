package main

import "fmt"

// Prosedur untuk mencetak deret bilangan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n) // Cetak nilai n
		if n%2 == 0 {        // Jika n genap
			n = n / 2
		} else { // Jika n ganjil
			n = 3*n + 1
		}
	}
	fmt.Println(1) // Cetak nilai akhir 1
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal deret: ")
	fmt.Scan(&n)  // Ambil input dari pengguna
	cetakDeret(n) // Panggil prosedur cetakDeret
}
