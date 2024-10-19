package main

import "fmt"

func main() {
	var a, b, c, d int // Mendeklarasikan variabel a, b, c, d menggunakan integer
	fmt.Print("Masukkan nilai: ") // Mencetak untuk inputan a, b, c, d
	fmt.Scan(&a, &b, &c, &d)

	var hasilPer, hasilKom int // Mendeklarasikan variabel hasil permutasi dan hasil kombinasi menggunakan integer

	// Menghitung permutasi dan kombinasi
	if a >= c && b>=d {  // Jika a >= c dan b >= d, maka hitung permutasi dan kombinasi

		// Permutasi untuk (a, c)
		permutation(a, c, &hasilPer)
		fmt.Println("Permutasi: ", hasilPer) // Cetak hasil permutasi (a, c)

		// Kombinasi untuk (a, c)
		kombination(a, c, &hasilKom)
		fmt.Println("Kombinasi: ", hasilKom) // Cetak hasil kombinasi (a, c)

		// Permutasi untuk (b, d)
		permutation(b, d, &hasilPer)
		fmt.Println("Permutasi: ", hasilPer) // Cetak hasil permutasi (b, d)

		// Kombinasi untuk (b, d)
		kombination(b, d, &hasilKom)
		fmt.Println("Kombinasi: ", hasilKom) // Cetak hasil kombinasi (b, d)

	} else { // Jika a >= c dan b >= d tidak terpenuhi, makan akan mencetak pesan
		fmt.Println("Tidak memenuhi kondisi") // Cetak pesan jika kondisi tidak terpenuhi
	}
}

// Prosedur untuk menghitung hasil faktorial 
func factorial(n int, hasil *int){
	*hasil = 1 // Inisialisasi hasil sebagai 1 (karena faktorial dimulai dari 1)
	for i := 1; i <= n; i++ {
		*hasil *= i 
	}
}

// Prosedur untuk menghitung hasil permutasi dari n, n-r
func permutation(n, r int, hasil *int) {
	var N, NR int // Inisialisasi variabel untuk menyimpan hasil faktorial n dan n-r
	factorial(n, &N) // Memanggil prosedur faktorial untuk menghitung faktorial dari n
	factorial(n-r, &NR) // Memanggil prosedur faktorial untuk menghitung faktorial dari n-r
	*hasil = N / NR // Menghitung permutasi

}

// Prosedur untuk menghitung hasil n, r, dan n-r 
func kombination(n, r int, hasil *int) {
	var N, R, NR int // Inisialisasi variabel untuk menyimpan hasil faktorial n, r, dan n-r
	factorial(n, &N) // Memanggil prosedur faktorial untuk menghitung faktorial dari n
	factorial(r, &R) // Memanggil prosedur faktorial untuk menghitung faktorial dari r
	factorial(n-r, &NR) // Memanggil prosedur faktorial untuk menghitung faktorial dari n-r
	*hasil = N / (R * NR) // Menghitung kombinasi
}
