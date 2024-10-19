//Haifa Zahra Azzimmi
//2311102163

package main

import (
	"fmt"
)

func cetakDeret(n int) {
	// Selama n belum bernilai 1, lanjutkan mencetak deret
	for n != 1 {
		// Cetak angka saat ini diikuti oleh spasi
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			// Jika angka genap, bagi dengan 2
			n /= 2
		} else {
			// Jika angka ganjil, kalikan dengan 3 dan tambahkan 1
			n = 3*n + 1
		}
	}
	// Terakhir, cetak angka 1
	fmt.Println(n)
}

func main() {
	// Deklarasi variabel untuk masukan
	var nilaiAwal int

	// Meminta input dari pengguna
	fmt.Print("Masukkan nilai suku awal: ")
	fmt.Scan(&nilaiAwal)

	// Panggil fungsi cetakDeret dengan nilai yang dimasukkan pengguna
	cetakDeret(nilaiAwal)
}
