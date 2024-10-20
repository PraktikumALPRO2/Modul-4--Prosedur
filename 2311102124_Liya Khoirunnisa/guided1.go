package main

import "fmt"

func main() {
	// Deklarasi variabel
	var a, b, hasil int

	// Meminta inputan dari pengguna
	fmt.Print("Masukkan dua angka: ")

	// Menyimpan inputan
	fmt.Scan(&a, &b)

	// Mengecek angka yang diinputkan pengguna
	if a >= b {
		// Jika hasil lebih besar atau sama dengan b
		permutasi(&a, &b, &hasil)

		// Cetak hasil
		fmt.Println("Hasil permutasi:", hasil)
	} else {
		// Jika a lebih kecil dari b
		permutasi(&b, &a, &hasil)

		// Cetak hasil
		fmt.Println("Hasil permutasi:", hasil)
	}
}

// Fungsi menghitung faktorial
func faktorial(n *int, hasil *int) {
	// Inisialisai
	*hasil = 1

	// Perulangan faktorial
	for i := 1; i <= *n; i++ {
		*hasil = *hasil * i
	}
}

// Fungsi menghitung permutasi
func permutasi(n, r, hasil *int) {
	var faktorialN, faktorialNR int

	// Menghitung faktorial dari n
	faktorial(n, &faktorialN)

	// Menghitung n-r
	nr := *n - *r

	// Menghitung faktorial dari n-r
	faktorial(&nr, &faktorialNR)

	// Hasil permutasi
	*hasil = faktorialN / faktorialNR
}
