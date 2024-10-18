// Caroline Carren
// 2311102174
// S1 IF 11 5

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("================================")
	fmt.Println("Program Permutasi & Kombinasi")
	fmt.Println("================================")

	var a, b, c, d int
	a, b, c, d = inputAngka()

	// Validasi input
	if !isValidInput(a, b, c, d) {
		fmt.Println("Input tidak valid. Pastikan a >= c dan b >= d.")
		return
	}

	// Hitung dan tampilkan hasil
	tampilkanHasil(a, b, c, d)
}

// Prosedur fungsi untuk menerima input angka dari pengguna
func inputAngka() (int, int, int, int) {
	var inputA, inputB, inputC, inputD string
	fmt.Print("Masukkan bilangan asli a: ")
	fmt.Scanln(&inputA)
	fmt.Print("Masukkan bilangan asli b: ")
	fmt.Scanln(&inputB)
	fmt.Print("Masukkan bilangan asli c: ")
	fmt.Scanln(&inputC)
	fmt.Print("Masukkan bilangan asli d: ")
	fmt.Scanln(&inputD)

	// Mengkonversi input string menjadi integer
	a, _ := strconv.Atoi(inputA)
	b, _ := strconv.Atoi(inputB)
	c, _ := strconv.Atoi(inputC)
	d, _ := strconv.Atoi(inputD)

	return a, b, c, d
}

// Prosedur untuk fungsi memvalidasi input
func isValidInput(a, b, c, d int) bool {
	return a >= c && b >= d
}

// Fungsi untuk menghitung faktorial dari n
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// Fungsi untuk menghitung permutasi P(n, r)
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// Fungsi untuk menghitung kombinasi C(n, r)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

// Fungsi untuk menampilkan hasil perhitungan
func tampilkanHasil(a, b, c, d int) {
	fmt.Println("--------------------------------")
	fmt.Println("Hasil Perhitungan:")
	fmt.Printf("P(%d, %d) = %d!/(%d-%d)! = %d\n", a, c, a, a, c, permutasi(a, c))
	fmt.Printf("C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", a, c, a, c, a, c, kombinasi(a, c))
	fmt.Printf("P(%d, %d) = %d!/(%d-%d)! = %d\n", b, d, b, b, d, permutasi(b, d))
	fmt.Printf("C(%d, %d) = %d!/((%d!)*(%d-%d)!) = %d\n", b, d, b, d, b, d, kombinasi(b, d))
	fmt.Println("--------------------------------")
}
