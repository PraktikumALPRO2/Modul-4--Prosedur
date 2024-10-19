package main

import (
	"fmt"
)

// Prosedur untuk mencetak deret bilangan
func cetakDeret(n int) {
	fmt.Print(n) // Mencetak suku pertama (nilai awal)

	for n != 1{ // Melanjutkan mencetak hingga suku bernilai 1
		if n%2 == 0 { // Jika n adalah bilangan genap, suku berikutnya adalah n dibagi 2
			n = n / 2
		} else { // Jika n adalah bilangan ganjil, suku berikutnya adalah dikalikan 3 lalu tambahkan 1
			n = 3*n + 1
		}
		fmt.Print(" ", n) // Mencetak suku berikutnya sampai suku terakhir bernilai 1
	}
	fmt.Print()
}

func main() {
	var n int
	fmt.Print("Masukkan nilai suku awal < 1000000: ") // Pengguna menginputkan nilai suku awal yang lebih kecil dari 1000000
	fmt.Scan(&n)

	if n >= 1 && n < 1000000 { // Menghitung atau memastikan nilai inputan valid
		cetakDeret(n) // Memanggil untuk mencetak deret menggunakan prosedur
	} else {
		fmt.Println("Masukkan harus berupa bilangan bulat positif < 1000000.") // Tampilan pesan jika inputan tidak valid
	}
}