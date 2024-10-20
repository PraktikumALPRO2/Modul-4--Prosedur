package main

import (
	"fmt"
)

// Prosedur untuk menghitung faktorial dan menyimpannya dalam pointer hasil
func hitungFaktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

// Prosedur untuk menghitung dan mencetak permutasi
func cetakPermutasi(n, r int) {
	if n < r {
		fmt.Println("Permutasi tidak dapat dihitung karena n < r.")
	} else {
		var faktN, faktNR int
		hitungFaktorial(n, &faktN)
		hitungFaktorial(n-r, &faktNR)
		perm := faktN / faktNR
		fmt.Printf("Permutasi (%dP%d) = %d\n", n, r, perm)
	}
}

// Prosedur untuk menghitung dan mencetak kombinasi
func cetakKombinasi(n, r int) {
	if n < r {
		fmt.Println("Kombinasi tidak dapat dihitung karena n < r.")
	} else {
		var faktN, faktR, faktNR int
		hitungFaktorial(n, &faktN)
		hitungFaktorial(r, &faktR)
		hitungFaktorial(n-r, &faktNR)
		komb := faktN / (faktR * faktNR)
		fmt.Printf("Kombinasi (%dC%d) = %d\n", n, r, komb)
	}
}

func main() {
	// Input empat bilangan: a, b, c, d
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan: ")
	fmt.Scan(&a, &b, &c, &d)
	// Baris pertama: permutasi dan kombinasi a terhadap c
	cetakPermutasi(a, c)
	cetakKombinasi(a, c)
	// Baris kedua: permutasi dan kombinasi b terhadap d
	cetakPermutasi(b, d)
	cetakKombinasi(b, d)
}
