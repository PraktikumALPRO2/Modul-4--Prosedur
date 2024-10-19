// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import "fmt"

// Fungsi untuk mencetak deret bilangan Skiena
func cetakDeret(n int) {
	// Looping hingga n mencapai 1
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2 // Jika n genap, bagi 2
		} else {
			n = 3*n + 1 // Jika n ganjil, kalikan 3 lalu tambah 1
		}
	}
	// Cetak nilai akhir (1)
	fmt.Println(n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&n)

	// Panggil prosedur cetakDeret untuk menampilkan hasilnya
	cetakDeret(n)
}
