package main

import (
	"fmt"
	"strings"
)

const maxTime = 301

// Prosedur untuk menghitung skor dan total waktu
func hitungSkor(soal [8]int, totalSoal *int, totalWaktu *int) {
	*totalSoal = 0
	*totalWaktu = 0
	for _, waktu := range soal {
		if waktu < maxTime {
			*totalSoal++
			*totalWaktu += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal [8]int
	var totalSoal, totalWaktu int
	maxSoal := -1
	minWaktu := maxTime * 8

	for {
		fmt.Print("Masukkan nama peserta (atau selesai jika ingin mengakhiri): ")
		fmt.Scan(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}

		fmt.Println("Masukkan waktu pengerjaan: ")
		for i := 0; i < 8; i++ {
			fmt.Scan(&soal[i])
		}

		// Hitung total soal yang diselesaikan dan total waktu yang dibutuhkan
		hitungSkor(soal, &totalSoal, &totalWaktu)

		if totalSoal > maxSoal || (totalSoal == maxSoal && totalWaktu < minWaktu) {
			maxSoal = totalSoal
			minWaktu = totalWaktu
			pemenang = nama
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}