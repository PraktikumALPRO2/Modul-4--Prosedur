package main

import (
	"fmt"
	"os"
)

// Fungsi untuk menghitung langkah Collatz
func hitungCollatz(n int) ([]int, int) {
	deret := []int{n}
	langkah := 0

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		deret = append(deret, n)
		langkah++
	}

	return deret, langkah
}

// Prosedur untuk mencetak deret
func cetakDeret(deret []int) {
	for i, nilai := range deret {
		if i == len(deret)-1 {
			fmt.Printf("%d", nilai)
		} else {
			fmt.Printf("%d ", nilai)
		}
	}
	fmt.Println()
}

// Fungsi untuk memvalidasi input
func validasiInput(n int) bool {
	return n > 0
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan positif: ")
	_, err := fmt.Scan(&n)

	if err != nil || !validasiInput(n) {
		fmt.Println("Input tidak valid. Mohon masukkan bilangan bulat positif.")
		os.Exit(1)
	}

	deret, langkah := hitungCollatz(n)

	fmt.Printf("\nDeret Collatz untuk %d:\n", n)
	cetakDeret(deret)
	fmt.Printf("Jumlah langkah: %d\n", langkah)
}
