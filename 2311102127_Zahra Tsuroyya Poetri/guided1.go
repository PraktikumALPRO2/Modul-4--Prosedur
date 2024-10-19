package main

import "fmt"

func main() {
	var a, b int // Mendeklarasikan variabel integer a dan b
	fmt.Print("Masukkan dua angka: ") // Mencetak pesan untuk pengguna dapat memasukkan nilai a dan b
	fmt.Scan(&a, &b) // Membaca input dari pengguna untuk dua nilai integer a dan b
	var hasil int // Mendeklarasikan variabel 'hasil' bertipe integer

	if a >= b { // Sebuah kondisi jika a lebih besar sama dengan b, maka:
		// Jika a lebih besar atau sama dengan b, panggil fungsi permutasi dengan parameter (a, b)
		permutasi(a, b, &hasil)
	} else { 
		// Jika b lebih besar dari a, panggil fungsi permutasi dengan parameter (b, a)
		permutasi(b, a, &hasil)
	}

	fmt.Println("Hasil permutasi: ", hasil) // Mencetak pesan hasil 
}


func faktorial(n int, hasil *int) { 
	*hasil = 1
	var i int
	// Loop untuk menghitung faktorial dari n
	for i = 1; i <= n; i++ {
		*hasil *= i // Mengalikan hasil dengan nilai i pada setiap iterasi
	}
	
}

func permutasi(n, r int, hasil *int) {
	// Menghitung permutasi nPr dengan membagi faktorial n dengan faktorial (n-r)
	var N, NR int
	faktorial(n, &N)
	faktorial(n-r, &NR)
	*hasil = N / NR
}

