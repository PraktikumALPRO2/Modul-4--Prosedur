/* Liya Khoirunnisa - 2311102124 */

package main

import "fmt"

func main() {
	// Deklarasi variabel
	var a, b, c, d, permutasiAC, kombinasiAC, permutasiBD, kombinasiBD int

	// Inputan dari pengguna
	fmt.Print("Masukkan angka a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan angka b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan angka c: ")
	fmt.Scan(&c)

	fmt.Print("Masukkan angka d: ")
	fmt.Scan(&d)

	// Mengecek apakah a >= c dan b >= d
	if a >= c && b >= d {
		// Menghitung permutasi dan kombinasi a terhadap c
		permutasi(&a, &c, &permutasiAC)
		kombinasi(&a, &c, &kombinasiAC)

		// Menghitung permutasi dan kombinasi b terhadap d
		permutasi(&b, &d, &permutasiBD)
		kombinasi(&b, &d, &kombinasiBD)

		// Cetak hasil untuk a terhadap c
		fmt.Println("Hasil permutasi a terhadap c:", permutasiAC)
		fmt.Println("Hasil kombinasi a terhadap c:", kombinasiAC)

		// Cetak hasil untuk b terhadap d
		fmt.Println("Hasil permutasi b terhadap d:", permutasiBD)
		fmt.Println("Hasil kombinasi b terhadap d:", kombinasiBD)
	} else {
		fmt.Println("Syarat a >= c dan b >= d harus terpenuhi.")
	}
}

// Fungsi menghitung faktorial
func faktorial(n *int, hasil *int) {
	// Inisialisasi
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

// Fungsi menghitung kombinasi
func kombinasi(n, r, hasil *int) {
	var faktorialN, faktorialR, faktorialNR int

	// Menghitung faktorial dari n
	faktorial(n, &faktorialN)

	// Menghitung faktorial dari r
	faktorial(r, &faktorialR)

	// Menghitung n-r
	nr := *n - *r

	// Menghitung faktorial dari n-r
	faktorial(&nr, &faktorialNR)

	// Hasil kombinasi
	*hasil = faktorialN / (faktorialR * faktorialNR)
}
