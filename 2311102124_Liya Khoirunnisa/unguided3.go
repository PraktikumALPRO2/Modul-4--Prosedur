/* Liya Khoirunnisa - 2311102124 */
package main

import "fmt"

// Prosedur untuk mencetak deret
func deretBilangan(n int) {
	fmt.Print("Deret bilangan: ")
	for n != 1 {
		// Cetak suku deret saat ini
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2 // Jika suku genap
		} else {
			n = 3*n + 1 // Jika suku ganjil
		}
	}

	// Cetak suku terakhir yang selalu bernilai 1
	fmt.Printf("1\n")
}

func main() {
	var n int
	fmt.Print("Masukkan angka (lebih kecil dari 1000000): ")
	fmt.Scan(&n)

	if n > 0 && n < 1000000 {
		deretBilangan(n) // Panggil prosedur deretBilangan dengan parameter n
	} else {
		fmt.Println("Input harus bilangan positif yang kurang dari 1000000.")
	}
}
